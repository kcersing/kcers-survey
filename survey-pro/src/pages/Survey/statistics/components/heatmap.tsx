import React, { useEffect, useRef } from 'react';

function isSupportCanvas() {
  const elem = document.createElement('canvas');
  return !!(elem.getContext && elem.getContext('2d'));
}

export const HeatMap = (props: { data: any[] }) => {
  const { data } = props;
  const mapRef = useRef<HTMLDivElement>(null);
  const mapInstance = useRef<any>(null);
  const heatmapInstance = useRef<any>(null);
  const scriptsLoaded = useRef(false);

  if (!isSupportCanvas()) {
    alert('热力图目前只支持有canvas支持的浏览器,您所使用的浏览器不能使用热力图功能~');
  }

  useEffect(() => {
    const loadScript = (src: string) => new Promise<void>((resolve, reject) => {
      if (document.querySelector(`script[src="${src}"]`)) {
        resolve();
        return;
      }
      const script = document.createElement('script');
      script.type = 'text/javascript';
      script.src = src;
      script.addEventListener('load', () => resolve());
      script.addEventListener('error', () => reject(new Error(`Failed to load ${src}`)));
      document.body.appendChild(script);
    });

    const createMap = () => {
      if (!mapRef.current) return;
      const T = (window as any).T;

      const map = new T.Map(mapRef.current);
      map.centerAndZoom(new T.LngLat(108.95, 34.27), 4);
      mapInstance.current = map;

      const heatmapOverlay = new T.HeatmapOverlay({
        radius: 30,
        maxOpacity: 0.9,
        minOpacity: 0.2,
        blur: 0.75,
        gradient: {
          0.0: 'blue',
          0.3: 'cyan',
          0.5: 'lime',
          0.7: 'yellow',
          0.85: 'orange',
          1.0: 'red',
        },
      });
      map.addOverLay(heatmapOverlay);
      heatmapInstance.current = heatmapOverlay;

      if (data && data.length > 0) {
        heatmapOverlay.setDataSet({ data, max: 300 });
        heatmapOverlay.show();
      }
    };

    if (scriptsLoaded.current) {
      // Scripts already loaded, just update data on existing map
      if (heatmapInstance.current && data && data.length > 0) {
        heatmapInstance.current.setDataSet({ data, max: 300 });
        heatmapInstance.current.show();
      }
      return;
    }

    scriptsLoaded.current = true;

    loadScript('https://api.tianditu.gov.cn/api?v=4.0&tk=516e46ec670dc4149ad67ed5020d99fd')
      .then(() => loadScript('/scripts/HeatmapOverlay.js'))
      .then(createMap)
      .catch(console.error);
  }, [data]);

  return (
    <div ref={mapRef} id="mapDiv" style={{ width: '100%', height: '900px' }}></div>
  );
};