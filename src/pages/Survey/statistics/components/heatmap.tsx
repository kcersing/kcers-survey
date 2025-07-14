import React, { useEffect, useRef } from 'react';

export const HeatMap = (props: { data: any }) => {
  const { data } = props;
  const mapRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const script = document.createElement('script');
    script.src = 'http://api.tianditu.gov.cn/api?v=4.0&tk=516e46ec670dc4149ad67ed5020d99fd';
    script.onload = () => {
      console.log('天地图 API 加载完成');
      if (mapRef.current) {
        try {
          // 使用 const 替代 var
          const map = new T.Map(mapRef.current);
          map.centerAndZoom(new T.LngLat(116.404, 39.915), 10);


          // 检查 T.HeatmapOverlay 是否可用
          if (typeof T.HeatmapOverlay === 'function') {
            // 使用 const 替代 var
            const heatmap = new T.HeatmapOverlay({
              radius: 40,
              max: 100
            });
            // 修正方法名拼写错误
            map.addOverlay(heatmap);
            // 使用 const 替代 var
            heatmap.setDataSet({ data, max: 100 });
            heatmap.show();
          } else {
            console.error('当前 API 版本不支持 T.HeatmapOverlay');
          }
        } catch (error) {
          console.error('地图初始化出错:', error);
        }
      }
    };
    script.onerror = () => {
      console.error('天地图 API 加载失败');
    };
    document.body.appendChild(script);

    return () => {
      document.body.removeChild(script);
    };
  }, []);

  return (
    <div ref={mapRef} style={{ width: '100%', height: '600px' }}></div>
  );
};
