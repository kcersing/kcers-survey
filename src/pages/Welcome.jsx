import { PageContainer } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import { Card, theme } from 'antd';
import React from 'react';
import { history } from '@@/core/history';
/**
 * 每个单独的卡片，为了复用样式抽成了组件
 * @param param0
 * @returns
 */
const InfoCard = ({ title, href, index, desc }) => {
    const { useToken } = theme;
    const { token } = useToken();
    return (<div style={{
            backgroundColor: token.colorBgContainer,
            boxShadow: token.boxShadow,
            borderRadius: '8px',
            fontSize: '14px',
            color: token.colorTextSecondary,
            lineHeight: '22px',
            padding: '16px 19px',
            minWidth: '220px',
            flex: 1,
        }}>
      <div style={{
            display: 'flex',
            gap: '4px',
            alignItems: 'center',
        }}>
        <div style={{
            width: 48,
            height: 48,
            lineHeight: '22px',
            backgroundSize: '100%',
            textAlign: 'center',
            padding: '8px 16px 16px 12px',
            color: '#FFF',
            fontWeight: 'bold',
            backgroundImage: "url('https://gw.alipayobjects.com/zos/bmw-prod/daaf8d50-8e6d-4251-905d-676a24ddfa12.svg')",
        }}>
          {index}
        </div>
        <div style={{
            fontSize: '16px',
            color: token.colorText,
            paddingBottom: 8,
        }}>
          {title}
        </div>
      </div>
      <div style={{
            fontSize: '14px',
            color: token.colorTextSecondary,
            textAlign: 'justify',
            lineHeight: '22px',
            marginBottom: 8,
        }}>
        {desc}
      </div>
      <a href={href} target="_blank" rel="noreferrer">
        了解更多 {'>'}
      </a>
    </div>);
};
const Welcome = () => {
    const { token } = theme.useToken();
    const { initialState } = useModel('@@initialState');
    const urlParams = new URL(window.location.href).searchParams;
    const usertoken = sessionStorage.getItem('token');
    console.log(usertoken);
    if (!usertoken) {
        history.push(urlParams.get('redirect') || '/user/login');
    }
    return (<PageContainer>
      <Card style={{
            borderRadius: 8,
        }}>
        <div style={{
            backgroundPosition: '100% -30%',
            backgroundRepeat: 'no-repeat',
            backgroundSize: '274px auto',
        }}>

          <p style={{
            fontSize: '14px',
            color: token.colorTextSecondary,
            lineHeight: '22px',
            marginTop: 16,
            marginBottom: 32,
            width: '65%',
        }}>
           疯狂敲代码中       </p>
          <div style={{
            display: 'flex',
            flexWrap: 'wrap',
            gap: 16,
        }}>
          </div>
        </div>
      </Card>
    </PageContainer>);
};
export default Welcome;
//# sourceMappingURL=Welcome.jsx.map