import { interviewerLogin } from '@/services/ant-design-pro/survey';
import { LockOutlined, MobileOutlined } from '@ant-design/icons';
import { LoginForm, ProFormText } from '@ant-design/pro-components';
import { history } from '@umijs/max';
import { message, Tabs } from 'antd';
import React, { useState } from 'react';

const InterviewerLogin: React.FC = () => {
  const [type, setType] = useState<string>('mobile');

  const handleSubmit = async (values: Record<string, any>) => {
    try {
      const res = await interviewerLogin({
        mobile: values.mobile,
        password: values.password || '',
      });
      if (res.code === 0) {
        if (res.data?.needSetup) {
          // 首次登录，跳转设置密码
          history.push(
            `/interviewer/setup-password?mobile=${encodeURIComponent(res.data.mobile)}&name=${encodeURIComponent(res.data.name)}`,
          );
          return;
        }
        sessionStorage.setItem('interviewer_token', res.data.token);
        sessionStorage.setItem('interviewer_mobile', res.data.mobile);
        message.success('登录成功');
        history.push('/interviewer/dashboard');
      } else {
        message.error(res.msg || '登录失败');
      }
    } catch (err: any) {
      message.error(err?.msg || '登录失败，请稍后重试');
    }
  };

  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        background: '#f0f2f5',
      }}
    >
      <div style={{ marginBottom: 24 }}>
        <h2 style={{ textAlign: 'center' }}>调查员登录</h2>
        <p style={{ textAlign: 'center', color: '#999' }}>请输入手机号，首次登录将引导设置密码</p>
      </div>
      <LoginForm
        contentStyle={{ minWidth: 320, maxWidth: 400 }}
        onFinish={handleSubmit}
        submitter={{
          searchConfig: { submitText: '登录' },
        }}
      >
        <Tabs
          activeKey={type}
          onChange={setType}
          centered
          items={[{ key: 'mobile', label: '手机号登录' }]}
        />
        <ProFormText
          name="mobile"
          fieldProps={{
            size: 'large',
            prefix: <MobileOutlined />,
            maxLength: 11,
          }}
          placeholder="请输入手机号"
          rules={[
            { required: true, message: '请输入手机号' },
            { pattern: /^1\d{10}$/, message: '请输入正确的手机号' },
          ]}
        />
        <ProFormText.Password
          name="password"
          fieldProps={{
            size: 'large',
            prefix: <LockOutlined />,
          }}
          placeholder="请输入密码（首次登录可不填）"
        />
      </LoginForm>
    </div>
  );
};

export default InterviewerLogin;