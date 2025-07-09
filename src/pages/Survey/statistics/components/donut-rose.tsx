import { Rose } from '@ant-design/plots';
import React from 'react';


export const DemoRose = () => {
  const config = {
    data: [

      { type: '分类一', value: 27 },
      { type: '分类二', value: 25 },
      { type: '分类三', value: 18 },
      { type: '分类四', value: 15 },
      { type: '分类五', value: 10 },
      { type: '其他', value: 5 },

    ],
    xField: 'type',
    yField: 'value',
    colorField: 'type',
    innerRadius: 0.2,
    state: {
      active: {
        fill: '#288AFF',
        stroke: 'black',
        lineWidth: 1,
        zIndex: 101,
      },
      inactive: {
        opacity: 0.5,
        zIndex: 100,
      },
    },
    legend: {
      color: {
        position: 'right',
        layout: {
          justifyContent: 'center',
        },
      },
    },
    interaction: {
      elementHighlight: true,
    },
    scale: { x: { padding: 0 } },
    axis: false,
    style: {
      lineWidth: 1,
      stroke: '#fff',
    },
    label: {
      text: 'type',
      fontSize: 16,
      fontWeight: 800,
      position: 'inside',
    },
  };
  return <Rose {...config} />;
};
