import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';

import {ProFormDependency, ProFormSelect, ProFormText,StepsForm,ProForm} from "@ant-design/pro-components";
import { queryCity, queryProvince } from '@/services/ant-design-pro/api';

import { Input } from 'antd';

const { TextArea } = Input;

const Address = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum,setCurrent } = props;
  const [value, setValue] = useState(0);

  const onChange = (e: RadioChangeEvent) => {
    setValue(e);
    addRespondent({
      surveyId:surveyId,
      type:"area",
      value:[e.value.toString()],
      sn:generateRandom,
    })
  };
  const onChange1 = (e) => {

    console.log(e)
    // setValue(e);
    addRespondent({
      surveyId:surveyId,
      type:"city",
      value:[e.toString()],
      sn:generateRandom,
    })
  };
  const onChange2 = (e: RadioChangeEvent) => {
    console.log(e)
    addRespondent({
      surveyId:surveyId,
      type:"city2",
      value:[e.toString()],
      sn:generateRandom,
    })
  };

  const onChange3 = (e) => {
    console.log(e)
    addRespondent({
      surveyId:surveyId,
      type:"address",
      value:[e.target.value.toString()],
      sn:generateRandom,
    })

  };


  return (
    <StepsForm.StepForm
      name={"StepsForm"}
      key={"StepsForm"}
      // onBlur={e => {   console.log(e.target.value)}}
    >
        <h3> 请输入您的所在的地址信息</h3>

        <ProFormSelect
          rules={[
            {
              required: true,
              message: '请输入您的所在省!',
            },
          ]}
          width="sm"
          fieldProps={{
            labelInValue: true,
          }}
          name="province"
          // className={styles.item}
          request={async () => {
            return queryProvince().then(({ data }) => {
              return data.map((item) => {
                return {
                  label: item.title,
                  value: item.value,
                };
              });
            });
          }}
          onChange={onChange}
        />
        <ProFormDependency name={['province']}>
          {({ province }) => {

            return (
              <ProFormSelect
                params={{
                  key: province?.value,
                }}
                name="city"
                onChange={onChange1}
                width="sm"
                rules={[
                  {
                    required: true,
                    message: '请输入您的所在城市!',
                  },
                ]}
                disabled={!province}
                // className={styles.item}
                request={async () => {
                  if (!province?.value) {
                    return [];
                  }
                  console.log(province)
                  return queryCity(province.value || '').then(({ data }) => {
                    return data.map((item) => {
                      return {
                        label: item.title,
                        value: item.value,
                      };
                    });
                  });
                }}
              />
            );
          }}
        </ProFormDependency>
      <ProFormDependency name={['province']}>
        {({ province }) => {

          return (
            <ProFormSelect
              onChange={onChange2}
              params={{
                key: province?.value,
              }}
              name="city2"
              width="sm"
              rules={[
                {
                  required: true,
                  message: '请输入您的所在区域!',
                },
              ]}
              disabled={!province}
              // className={styles.item}
              request={async () => {
                if (!province?.value) {
                  return [];
                }

                return queryCity(province.value || '').then(({ data }) => {
                  return data.map((item) => {
                    return {
                      label: item.title,
                      value: item.value,
                    };
                  });
                });
              }}
            />
          );
        }}
      </ProFormDependency>


        <TextArea
          style={{ width: '60%' }}

          name="address"
          label="街道地址"
          onChange={onChange3}
        />

    </StepsForm.StepForm>
      );
};

export default Address;
