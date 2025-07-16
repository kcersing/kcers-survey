import React, { useEffect, useRef } from 'react';

export const HeatMap = (props: { data: any }) => {
  const { data } = props;
  const mapRef = useRef<HTMLDivElement>(null);


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

console.log(1)

    loadHeat().then(() => {
      console.log(2)
      loadHeatmap().then(() => {
        console.log(3)
        if (mapRef.current) {
          try {
            const T = window.T
            const map = new T.Map(mapRef.current);
            map.centerAndZoom(new T.LngLat(116.404, 39.915), 10);

              const heatmap = new T.HeatmapOverlay({
                radius: 40,
                max: 100
              });

              map.addOverLay(heatmap);
              heatmap.setDataSet({ data, max: 100 });
              heatmap.show();

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
