import { Column } from '@ant-design/plots';
import React from 'react';

export const DemoCustomColor = () => {
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
    scale: {
      color: {
        range: ['#f4664a', '#faad14', '#a0d911', '#52c41a', '#13c2c2', '#1890ff', '#2f54eb', '#722ed1'],
      },
    },
  };
  return <Column {...config} />;
};
