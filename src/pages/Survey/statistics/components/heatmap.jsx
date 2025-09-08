import React, { useRef } from 'react';
function isSupportCanvas() {
    const elem = document.createElement('canvas');
    return !!(elem.getContext && elem.getContext('2d'));
}
export const HeatMap = (props) => {
    const { data } = props;
    const mapRef = useRef(null);
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
        script.src = "/scripts/HeatmapOverlay.js";
        script.addEventListener('load', () => {
            resolve('loaded');
        });
        document.body.appendChild(script);
    });
    function onLoad() {
        loadHeat().then(() => {
            loadHeatmap().then(() => {
                console.log('地图初始化');
                const T = window.T;
                // const map = new T.Map('mapDiv');
                const map = new T.Map(mapRef.current);
                map.centerAndZoom(new T.LngLat(108.95, 34.27), 4);
                const heatmapOverlay = new T.HeatmapOverlay({
                    "radius": 30,
                });
                map.addOverLay(heatmapOverlay);
                heatmapOverlay.setDataSet({ data: data, max: 300 });
                heatmapOverlay.show();
            });
        });
    }
    onLoad();
    return (<div ref={mapRef} id="mapDiv" style={{ width: '100%', height: '900px' }}></div>);
};
//# sourceMappingURL=heatmap.jsx.map