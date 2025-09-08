import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';

import {ProFormDependency, ProFormSelect,ProFormUploadDragger, ProFormUploadButton,ProFormText,StepsForm,ProForm} from "@ant-design/pro-components";
import { pubUpload, queryCity, queryProvince } from '@/services/api';

import { Input, Form,Checkbox,message } from 'antd';

const SFUpload = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum,setCurrent } = props;

  return (

    <StepsForm.StepForm
      name="StepsFormUpload"
      key="StepsFormUpload"
    // onBlur={e => {   console.log(e.target.value)}}
    >

      <Form.Item>

        <ProFormUploadButton
          name="上传合照"
          label="上传合照"
          max={10}
          action={(file)=>{
            pubUpload({file}).then((res)=>{
              console.log(res)
              if (res.code === 0) {
                message.success(`上传成功`);
                addRespondent({
                  surveyId:surveyId,
                  type:"image",
                  value:[res.data.url],
                  sn:generateRandom,
                })
              }
            })
          }}
          listType="picture-card"
          // 限制上传文件类型为图片文件
          accept="image/*"
          // 限制文件大小为 5MB
          maxSize={5 * 1024}
        />
      </Form.Item>
      </StepsForm.StepForm>
);
};

export default SFUpload;
