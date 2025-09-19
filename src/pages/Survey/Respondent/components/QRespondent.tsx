

import React, { type ReactElement, useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox } from 'antd';
import {ProFormText, ProFormTextArea, StepsForm} from "@ant-design/pro-components";
import { RuleType, StoreValue } from 'rc-field-form/lib/interface';

const style: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
};

const QRespondent = (props) => {

  const { surveyId,questions, generateRandom, addRespondent, setCurrentNum ,setCurrent} = props;
  setCurrentNum(0)
  return (
    <StepsForm.StepForm
      name={`key_${questions.length+1}`}
      key={`key_${questions.length+1}`}
  // onBlur={e => {   console.log(e.target.value)}}
>
  <ProFormText width="md"
  onChange={(e)=>{
    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'respondent',
      value:[e.target.value],
      sn:generateRandom,
    })
  }}
  label="访谈人姓名" rules={[{ required: true }]} name={'respondent'}
  />
  <ProFormText width="md"  onChange={(e)=>{

    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'respondentPhone',
      value:[e.target.value],
      sn:generateRandom,
    })}}
               label="联系电话"
               rules={[{ required: true,  len:11}]} name={'respondentPhone'}

  />
  <ProFormText width="md"  onChange={(e)=>{

    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'researcher',
      value:[e.target.value],
      sn:generateRandom,
    }) }} label="调研员姓名" rules={[{ required: true }]} name={'researcher'}
  />
  <ProFormText length={11}  width="md"  onChange={(e)=>{

    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'researcherPhone',
      value:[e.target.value],
      sn:generateRandom,
    })}} label="联系电话"  rules={[{ required: true, len:11 }]} name={'researcherPhone'}
  />
  </StepsForm.StepForm>
);
};

export default QRespondent;
