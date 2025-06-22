import React, { useState, useEffect } from 'react';
import { ProForm, ProFormText, ProFormSelect, ProFormDatePicker,  } from '@ant-design/pro-components';
import { history } from 'umi';
import { createSurvey, updateSurvey, getSurvey } from '@/services/ant-design-pro/survey';

import {Button, message  } from 'antd';

const SurveyEdit: React.FC<{ match: { params: { id?: string } } }> = ({ match }) => {
  // const surveyId = match.params.id;
  const surveyId = "1";
  const isEditMode = !!surveyId;
  const [form] = ProForm.useForm();

  useEffect(() => {
    if (isEditMode) {
      getSurvey({'surveyId':parseInt(surveyId)}).then(data => {
        form.setFieldsValue(data);
      });
    }
  }, []);

  const onFinish = async (values: any) => {
    try {
      if (isEditMode) {
        await updateSurvey({'surveyId':parseInt(surveyId)}, values);
        message.success('问卷更新成功');
      } else {
        await createSurvey(values);
        message.success('问卷创建成功');
      }
      history.push('/survey/list');
    } catch (error) {
      message.error(isEditMode ? '问卷更新失败' : '问卷创建失败');
    }
  };

  return (
    <ProForm
      form={form}
      layout="vertical"
      onFinish={onFinish}
      initialValues={{
        status: 'draft',
      }}
    >
      <ProFormText
        name="title"
        label="标题"
        rules={[{ required: true, message: '请输入问卷标题' }]}
      />

      <ProFormText
        name="description"
        label="描述"
        textarea
        rows={4}
      />

      <ProFormSelect
        name="status"
        label="状态"
        options={[
          { label: '草稿', value: 'draft' },
          { label: '发布中', value: 'active' },
          { label: '已关闭', value: 'closed' },
        ]}
      />

      <ProFormDatePicker
        name="start_time"
        label="开始时间"
        showTime
      />

      <ProFormDatePicker
        name="end_time"
        label="结束时间"
        showTime
      />

      <div style={{ marginTop: 24, display: 'flex', justifyContent: 'flex-end' }}>
        <Button onClick={() => history.goBack()} style={{ marginRight: 8 }}>
          返回
        </Button>
        <Button type="primary" htmlType="submit">
          保存
        </Button>
      </div>
    </ProForm>
  );
};

export default SurveyEdit;
