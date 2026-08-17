import { interviewerChangePassword } from '@/services/ant-design-pro/survey';
import { LockOutlined } from '@ant-design/icons';
import { PageContainer, ProFormText } from '@ant-design/pro-components';
import { history } from '@umijs/max';
import { Button, Card, Form, message } from 'antd';
import React, { useEffect } from 'react';

const ChangePassword: React.FC = () => {
  const [form] = Form.useForm();
  const token = sessionStorage.getItem('interviewer_token') || '';

  useEffect(() => {
    if (!token) {
      history.push('/interviewer/login');
    }
  }, []);

  const handleSubmit = async (values: Record<string, any>) => {
    if (values.newPassword !== values.confirmPassword) {
      message.error('两次输入的新密码不一致');
      return;
    }
    try {
      const res = await interviewerChangePassword(
        { oldPassword: values.oldPassword, newPassword: values.newPassword },
        token,
      );
      if (res.code === 0) {
        message.success('密码修改成功');
        form.resetFields();
      } else {
        message.error(res.msg || '修改失败');
      }
    } catch {
      message.error('修改失败，请稍后重试');
    }
  };

  return (
    <PageContainer
      header={{ title: '修改密码', onBack: () => history.push('/interviewer/dashboard') }}
    >
      <div style={{ maxWidth: 500, margin: '0 auto', padding: '40px 0' }}>
        <Card>
          <Form form={form} layout="vertical" onFinish={handleSubmit}>
            <ProFormText.Password
              name="oldPassword"
              label="原密码"
              fieldProps={{ prefix: <LockOutlined />, size: 'large' }}
              rules={[{ required: true, message: '请输入原密码' }]}
            />
            <ProFormText.Password
              name="newPassword"
              label="新密码"
              fieldProps={{ prefix: <LockOutlined />, size: 'large' }}
              rules={[
                { required: true, message: '请输入新密码' },
                { min: 6, message: '密码至少6位' },
              ]}
            />
            <ProFormText.Password
              name="confirmPassword"
              label="确认新密码"
              fieldProps={{ prefix: <LockOutlined />, size: 'large' }}
              rules={[
                { required: true, message: '请再次输入新密码' },
                { min: 6, message: '密码至少6位' },
              ]}
            />
            <Form.Item>
              <Button type="primary" htmlType="submit" block size="large">
                修改密码
              </Button>
            </Form.Item>
          </Form>
        </Card>
      </div>
    </PageContainer>
  );
};

export default ChangePassword;
