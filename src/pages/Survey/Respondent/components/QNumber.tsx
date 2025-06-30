

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox } from 'antd';
import {ProFormDigit, ProFormTextArea} from "@ant-design/pro-components";



const QNumber = (props) => {

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
        console.log(String(e))
        console.log(String(e) === jumpRule.answer)
        if (String(e) === jumpRule.answer) {
          setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);

        }
      }
    }
  };

  return(
  <Form.Item name={['question', "'"+question.id+"'"]}   required={question.required===1} >
    <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
  <ProFormDigit
    width="md"
    placeholder="请输入数字"
    name={['question', question.id]}
    style={{Width: 60}}
    onChange={onChange}
    rules={[{required: question.required === 1, message: '必填项'}]}
  />    </Form.Item>);
};

export default QNumber;
