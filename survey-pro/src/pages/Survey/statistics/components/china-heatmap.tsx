import { Button } from 'antd';
import * as echarts from 'echarts';
import React, { useEffect, useRef } from 'react';
import { computeHeatMax, toHeatmapSeriesData } from './heatmap-utils';

const CHINA_CENTER: [number, number] = [104.19, 35.86];

export const ChinaHeatMap = (props: { data: any[] }) => {
  const { data } = props;
  const chartRef = useRef<HTMLDivElement>(null);
  const chartInstance = useRef<any>(null);
  const geoJsonRef = useRef<any>(null);

  const resetView = () => {
    const chart = chartInstance.current;
    if (!chart) return;
    chart.setOption({ geo: { zoom: 1, center: CHINA_CENTER } });
  };

  useEffect(() => {
    const render = async () => {
      if (!chartRef.current) return;
      if (!geoJsonRef.current) {
        const res = await fetch('/maps/china.json');
        geoJsonRef.current = await res.json();
        echarts.registerMap('china', geoJsonRef.current);
      }
      if (!chartInstance.current) {
        chartInstance.current = echarts.init(chartRef.current);
      }
      const chart = chartInstance.current;
      chart.setOption({
        tooltip: {
          formatter: (params: any) => {
            const v = params.value as number[];
            return `调研点数量：${v[2]}`;
          },
        },
        visualMap: {
          min: 0,
          max: computeHeatMax(data),
          calculable: true,
          orient: 'vertical',
          left: 10,
          bottom: 60,
          inRange: {
            color: ['#313695', '#74add1', '#abd9e9', '#fee090', '#fdae61', '#f46d43', '#d73027'],
          },
        },
        geo: {
          map: 'china',
          roam: true,
          zoom: 1,
          center: CHINA_CENTER,
          label: {
            show: true,
            fontSize: 9,
            color: '#555',
          },
          itemStyle: {
            areaColor: '#f2f4f7',
            borderColor: '#aab4c0',
          },
        },
        series: [
          {
            name: '调研点',
            type: 'heatmap',
            coordinateSystem: 'geo',
            data: toHeatmapSeriesData(data),
            pointSize: 8,
            blurSize: 12,
          },
        ],
      });
    };
    render().catch(console.error);

    const onResize = () => {
      chartInstance.current?.resize();
    };
    window.addEventListener('resize', onResize);
    return () => {
      window.removeEventListener('resize', onResize);
      chartInstance.current?.dispose();
      chartInstance.current = null;
    };
  }, [data]);

  return (
    <div style={{ position: 'relative', width: '100%', height: '100%' }}>
      <div ref={chartRef} style={{ width: '100%', height: '100%' }} />
      <Button
        size="small"
        onClick={resetView}
        style={{ position: 'absolute', top: 8, right: 8, zIndex: 10 }}
      >
        回到全貌
      </Button>
    </div>
  );
};