import { Base } from '@ant-design/plots';
import React from 'react';


export const Demobase = (props: { data: any }) => {
  const { data } = props;
  const config = {
    type: 'spaceLayer',
    data: {
      type: 'fetch',
      value: data,
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

