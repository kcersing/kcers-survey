import { Column } from '@ant-design/plots';
import React from 'react';
export const DemoCustomColor = (props) => {
    console.log("DemoCustomColor");
    const { data } = props;
    if (!data) {
        return null;
    }
    const config = {
        data: data,
        xField: 'type',
        yField: 'value',
        colorField: 'type',
        scale: {
            color: {
                range: ['#f4664a', '#faad14', '#a0d911', '#52c41a', '#13c2c2', '#1890ff', '#2f54eb', '#722ed1'],
            },
        },
    };
    return <Column {...config}/>;
};
//# sourceMappingURL=custom-color.jsx.map