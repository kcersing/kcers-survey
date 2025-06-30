

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, DatePicker,Form,Checkbox } from 'antd';
import {ProFormDatePicker, ProFormRate, ProFormTextArea} from '@ant-design/pro-components';
import dayjs from 'dayjs';


const QDate = (props) => {

  const { surveyId, question, generateRandom, addRespondent, setCurrentNum } = props;

  if (!question ){return null}
  const onChange = (e: RadioChangeEvent) => {

    console.log(e)

    addRespondent({
      surveyId:surveyId,
      questionId:question.id,
      type:question.type,
      value:e,
      sn:generateRandom,
    })
    if (question.jumpRules) {
      for (const jumpRule of question.jumpRules) {
        if (String(e) === jumpRule.answer) {
          setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);
        }
      }
    }
  };


  return (
    <Form.Item  name={['question', "'"+question.id+"'"]}   required={question.required===1} >
      <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
       <DatePicker
      width="md" label={question.content}
      name={['question', question.id]}
      placeholder="请选择日期"
      defaultValue={dayjs('1965-01-01', 'YYYY-MM-DD')}
      defaultPickerValue={dayjs('1965-01-01', 'YYYY-MM-DD')}
      onChange={onChange}
      rules={[{required: question.required === 1, message: '必填项'}]}/>;
    </Form.Item>
  );
};

export default QDate;
