import { Dendrogram, G6 } from '@ant-design/graphs';
import React, { useEffect, useState } from 'react';
const { treeToGraphData } = G6;
export const DemoDendrogram = (props) => {
    const { data } = props;
    const [data, setData] = useState(undefined);
    useEffect(() => {
        fetch('https://gw.alipayobjects.com/os/antvdemo/assets/data/algorithm-category.json')
            .then((res) => res.json())
            .then((data) => setData(treeToGraphData(data)));
    }, []);
    const options = {
        autoFit: 'view',
        data,
        direction: 'vertical',
        compact: true,
    };
    return <Dendrogram {...options}/>;
};
//# sourceMappingURL=vertical-tidy-tree.jsx.map