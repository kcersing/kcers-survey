import { Base } from '@ant-design/plots';
import React from 'react';


export const Demobase = () => {
  const config = {
    type: 'spaceLayer',
    data: {
      type: 'fetch',
      value: [
        { type: '分类一', value: 27 },
        { type: '分类二', value: 25 },
        { type: '分类三', value: 18 },
        { type: '分类四', value: 15 },
        { type: '分类五', value: 10 },
        { type: '其他', value: 5 },],
    },
    children: [
      {
        type: 'interval',
        encode: { x: 'type', y: 'value', color: 'type' },
        transform: [{ type: 'sortX', reverse: true, by: 'y' }],
        scale: { color: { palette: 'cool', offset: (t) => t * 0.8 + 0.1 } },
      },
      {
        type: 'interval',
        x: 300,
        y: 50,
        width: 300,
        height: 300,
        encode: { y: 'value', color: 'type' },
        transform: [{ type: 'stackY' }],
        scale: { color: { palette: 'cool', offset: (t) => t * 0.8 + 0.1 } },
        coordinate: { type: 'theta' },
        legend: false,
      },
    ],
  };
  return <Base {...config} />;
};

