import { Button } from 'antd';
import React, { useEffect, useRef } from 'react';
import { computeHeatMax, parsePoints } from './heatmap-utils';

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

  const fitToData = () => {
    const map = mapInstance.current;
    const T = (window as any).T;
    if (!map || !T) return;
    const points = parsePoints(data);
    if (points.length === 0) {
      map.centerAndZoom(new T.LngLat(108.95, 34.27), 4);
      return;
    }
    if (points.length === 1) {
      map.centerAndZoom(new T.LngLat(points[0].lng, points[0].lat), 12);
      return;
    }
    map.setViewport(points.map((p) => new T.LngLat(p.lng, p.lat)));
  };

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
      mapInstance.current = map;
      map.centerAndZoom(new T.LngLat(108.95, 34.27), 4);

      const heatmapOverlay = new T.HeatmapOverlay({
        radius: 12,
        maxOpacity: 0.9,
        minOpacity: 0.2,
        blur: 0.3,
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
        heatmapOverlay.setDataSet({ data, max: computeHeatMax(data) });
        heatmapOverlay.show();
      }
      fitToData();
    };

    if (scriptsLoaded.current) {
      // Scripts already loaded, just update data on existing map
      if (heatmapInstance.current && data && data.length > 0) {
        heatmapInstance.current.setDataSet({ data, max: computeHeatMax(data) });
        heatmapInstance.current.show();
      }
      return;
    }

    scriptsLoaded.current = true;

    loadScript('https://api.tianditu.gov.cn/api?v=4.0&tk=516e46ec670dc4149ad67ed5020d99fd')
      .then(() => loadScript('/scripts/HeatmapOverlay.js'))
      .then(createMap)
      .catch(console.error);

    const onResize = () => {
      mapInstance.current?.checkResize();
    };
    window.addEventListener('resize', onResize);
    return () => window.removeEventListener('resize', onResize);
  }, [data]);

  return (
    <div style={{ position: 'relative', width: '100%', height: '100%' }}>
      <div ref={mapRef} id="mapDiv" style={{ width: '100%', height: '100%' }} />
      <Button
        size="small"
        onClick={fitToData}
        style={{ position: 'absolute', top: 8, right: 8, zIndex: 10 }}
      >
        回到全貌
      </Button>
    </div>
  );
};