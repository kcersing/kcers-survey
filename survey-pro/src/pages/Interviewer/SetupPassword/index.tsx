import { interviewerSendSms, interviewerSetupPassword } from '@/services/ant-design-pro/survey';
import { LockOutlined, SafetyCertificateOutlined } from '@ant-design/icons';
import { ProFormText } from '@ant-design/pro-components';
import { history, useSearchParams } from '@umijs/max';
import { Button, Card, Form, Input, message } from 'antd';
import React, { useState } from 'react';

const SetupPassword: React.FC = () => {
  const [searchParams] = useSearchParams();
  const mobile = searchParams.get('mobile') || '';
  const name = searchParams.get('name') || '';
  const [form] = Form.useForm();
  const [sending, setSending] = useState(false);
  const [countdown, setCountdown] = useState(0);

  if (!mobile) {
    history.push('/interviewer/login');
    return null;
  }

  const handleSendSms = async () => {
    setSending(true);
    try {
      const res = await interviewerSendSms({ mobile });
      if (res.code === 0) {
        message.success('验证码已发送');
        setCountdown(60);
        const timer = setInterval(() => {
          setCountdown((prev) => {
            if (prev <= 1) {
              clearInterval(timer);
              return 0;
            }
            return prev - 1;
          });
        }, 1000);
      } else {
        message.error(res.msg || '发送失败');
      }
    } catch {
      message.error('发送失败，请稍后重试');
    } finally {
      setSending(false);
    }
  };

  const handleSubmit = async (values: Record<string, any>) => {
    if (values.password !== values.confirmPassword) {
      message.error('两次输入的密码不一致');
      return;
    }
    try {
      const res = await interviewerSetupPassword({
        mobile,
        name,
        password: values.password,
        smsCode: values.smsCode,
      });
      if (res.code === 0) {
        sessionStorage.setItem('interviewer_token', res.data.token);
        sessionStorage.setItem('interviewer_mobile', res.data.mobile);
        message.success('密码设置成功，已自动登录');
        history.push('/interviewer/dashboard');
      } else {
        message.error(res.msg || '设置失败');
      }
    } catch {
      message.error('设置失败，请稍后重试');
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
        <h2 style={{ textAlign: 'center' }}>设置登录密码</h2>
        <p style={{ textAlign: 'center', color: '#999' }}>
          {name}（{mobile}）
        </p>
        <p style={{ textAlign: 'center', color: '#666' }}>
          首次登录，请验证手机号并设置密码
        </p>
      </div>
      <Card style={{ width: 400 }}>
        <Form form={form} layout="vertical" onFinish={handleSubmit}>
          <Form.Item label="短信验证码" required>
            <div style={{ display: 'flex', gap: 12 }}>
              <Form.Item
                name="smsCode"
                noStyle
                rules={[{ required: true, message: '请输入验证码' }]}
              >
                <Input
                  prefix={<SafetyCertificateOutlined />}
                  placeholder="请输入验证码"
                  size="large"
                  style={{ flex: 1 }}
                  maxLength={6}
                />
              </Form.Item>
              <Button
                size="large"
                disabled={countdown > 0}
                loading={sending}
                onClick={handleSendSms}
                style={{ minWidth: 110 }}
              >
                {countdown > 0 ? `${countdown}s` : '获取验证码'}
              </Button>
            </div>
          </Form.Item>
          <ProFormText.Password
            name="password"
            label="密码"
            fieldProps={{ prefix: <LockOutlined />, size: 'large' }}
            rules={[
              { required: true, message: '请设置密码' },
              { min: 6, message: '密码至少6位' },
            ]}
          />
          <ProFormText.Password
            name="confirmPassword"
            label="确认密码"
            fieldProps={{ prefix: <LockOutlined />, size: 'large' }}
            rules={[
              { required: true, message: '请再次输入密码' },
              { min: 6, message: '密码至少6位' },
            ]}
          />
          <Form.Item>
            <Button type="primary" htmlType="submit" block size="large">
              设置密码并登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default SetupPassword;