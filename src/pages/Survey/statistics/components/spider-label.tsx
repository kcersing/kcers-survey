import { Pie } from '@ant-design/plots';
import React from 'react';


export const DemoPie = (props: { data: any }) => {
  const { data } = props;
  const config = {
    data,
    angleField: 'value',
    colorField: 'type',
    radius: 0.8,
    label: {
      text: (d) => `${d.type}\n ${d.value}`,
      position: 'spider',
    },
    legend: {
      color: {
        title: false,
        position: 'right',
        rowPadding: 5,
      },
    },
  };
  return <Pie {...config} />;
};

