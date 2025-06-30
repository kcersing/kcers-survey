

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox } from 'antd';
import {ProFormText, ProFormTextArea, StepsForm} from "@ant-design/pro-components";

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
      value:e,
      sn:generateRandom,
    })
  }}
  label="访谈人姓名" rules={[{ required: true, message: '必填项' }]} name={'respondent'} />
  <ProFormText width="md"  onChange={(e)=>{
    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'respondentPhone',
      value:e,
      sn:generateRandom,
    })}}
  label="联系电话" rules={[{ required: true, message: '必填项' }]} name={'respondentPhone'} />
  <ProFormText width="md"  onChange={(e)=>{
    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'researcher',
      value:e,
      sn:generateRandom,
    }) }} label="调研员姓名" rules={[{ required: true, message: '必填项' }]} name={'researcher'} />
  <ProFormText width="md"  onChange={(e)=>{
    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'researcherPhone',
      value:e,
      sn:generateRandom,
    })}} label="联系电话"   rules={[{ required: true, message: '必填项' }]} name={'researcherPhone'} />
  <ProFormText  width="md"  onChange={(e)=>{
    addRespondent({
      surveyId:surveyId,
      questionId:0,
      type:'ditu',
      value:e,
      sn:generateRandom,
    })}} label="地图插件预留" name={'ditu'} />
  </StepsForm.StepForm>
);
};

export default QRespondent;
