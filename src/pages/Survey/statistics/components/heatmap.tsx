import React, { useEffect, useRef } from 'react';
function isSupportCanvas() {
  var elem = document.createElement('canvas');
  return !!(elem.getContext && elem.getContext('2d'));
}
export const HeatMap = (props: { data: any }) => {
  const { data } = props;
  const mapRef = useRef<HTMLDivElement>(null);

  if (!isSupportCanvas()) {
    alert('热力图目前只支持有canvas支持的浏览器,您所使用的浏览器不能使用热力图功能~');
  }
  const loadHeat = () => new Promise((resolve) => {
    const script = document.createElement('script');
    script.type = 'text/javascript';
    script.src = 'https://api.tianditu.gov.cn/api?v=4.0&tk=516e46ec670dc4149ad67ed5020d99fd';
    script.addEventListener('load', () => {
      resolve('loaded');
    });
    document.body.appendChild(script);
  });
  const loadHeatmap = () => new Promise((resolve) => {
    const script = document.createElement('script');
    script.type = 'text/javascript';
    script.src = 'https://survey.367281.com/scripts/HeatmapOverlay.js';
    script.addEventListener('load', () => {
      resolve('loaded');
    });
    document.body.appendChild(script);
  });


  useEffect(() => {
    loadHeat().then(() => {
      loadHeatmap().then(() => {

        if (mapRef.current) {
          try {
            const T = window.T
            const map = new T.Map(mapRef.current);
            map.centerAndZoom(new T.LngLat(116.404, 39.915), 4);

              const heatmap = new T.HeatmapOverlay({
                radius: 30
              });

              map.addOverLay(heatmap);
              heatmap.setDataSet({ data, max: 500 });
              heatmap.show();

              const heatmapCanvas = mapRef.current?.querySelector('canvas');
              if (heatmapCanvas) {
                heatmapCanvas.style.pointerEvents = 'none';
              }

          } catch (error) {
            console.error('地图初始化出错:', error);
          }
        }
      });
    });
  }, [data]);

  return (
    <div ref={mapRef} style={{ width: '100%', height: '600px' }}></div>
  );
};
