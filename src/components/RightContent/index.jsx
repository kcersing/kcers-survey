import { QuestionCircleOutlined } from '@ant-design/icons';
import { SelectLang as UmiSelectLang } from '@umijs/max';
import React from 'react';
export const SelectLang = () => {
    return (<UmiSelectLang style={{
            padding: 4,
        }}/>);
};
export const Question = () => {
    return (<div style={{
            display: 'flex',
            height: 26,
        }} onClick={() => {
            window.open('http://survey.367281.com/');
        }}>
      <QuestionCircleOutlined />
    </div>);
};
//# sourceMappingURL=index.jsx.map