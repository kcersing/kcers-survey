import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';

import {ProFormDependency, ProFormSelect, ProFormText,StepsForm,ProForm} from "@ant-design/pro-components";
import { queryCity, queryProvince } from '@/services/ant-design-pro/api';

import { Input ,Form} from 'antd';

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
  const onChangeCity = (e) => {

    console.log(e)
    // setValue(e);
    addRespondent({
      surveyId:surveyId,
      type:"city",
      value:[e.toString()],
      sn:generateRandom,
    })
  };
  const onChangeDistrict = (e: RadioChangeEvent) => {
    console.log(e)
    addRespondent({
      surveyId:surveyId,
      type:"district",
      value:[e.toString()],
      sn:generateRandom,
    })
  };

  const onChangeAddress = (e) => {
    console.log(e)
    addRespondent({
      surveyId:surveyId,
      type:"address",
      value:[e.target.value.toString()],
      sn:generateRandom,
    })

  };

  const onChangeVillage = (e: RadioChangeEvent) => {
    console.log(e)
    addRespondent({
      surveyId:surveyId,
      type:"village",
      value:[e.toString()],
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
          label={'省'}
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
                label={'市（州）'}
                params={{
                  key: province?.value,
                }}
                name="city"
                onChange={onChangeCity}
                width="sm"
                rules={[
                  {
                    required: true,
                    message: '请输入您的所在市（州）',
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
      {/* 区县选择框 */}
      <ProFormDependency name={['city']}>
        {({ city }) => {
          return (
            <ProFormSelect
              onChange={onChangeDistrict}
              label={'县（区、旗）'}
              params={{
                city: city,
              }}
              name="district"
              width="sm"
              disabled={!city}
              // className={styles.item}
              request={async () => {
                if ( !city) {
                  return [];
                }

                return queryCity(city || '').then(({ data }) => {
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
      <ProFormDependency name={['district',]}>
        {({ district }) => {

          return (
            <ProFormSelect
              label={'乡（镇）'}
              onChange={onChangeVillage}
              params={{
                key: district,
              }}
              name="village"
              width="sm"
              disabled={!district}
              // className={styles.item}
              request={async () => {
                if (!district) {
                  return [];
                }

                return queryCity(district || '').then(({ data }) => {
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
      <Form.Item label={'村'} >
        <TextArea
          style={{ width: '60%' }}
          name="address"
          onChange={onChangeAddress}
        />
      </Form.Item>
    </StepsForm.StepForm>
      );
};

export default Address;
