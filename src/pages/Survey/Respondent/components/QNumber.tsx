import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox } from 'antd';
import {ProFormDigit, ProFormTextArea} from "@ant-design/pro-components";
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';

const QNumber = (props) => {

  const { surveyId, question, generateRandom, addRespondent,setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState(0);
  if (!question ){return null}
  const onChange = (e: RadioChangeEvent) => {
    console.log(e)
    setValue(e);
    addRespondent({
      surveyId:surveyId,
      questionId:question.id,
      type:question.type,
      value:[e.toString()],
      sn:generateRandom,
    })
    if (question.jumpRules) {
      for (const jumpRule of question.jumpRules) {
        console.log(String(e))
        console.log(jumpRule.operators === 'equals' && String(e) === jumpRule.answer)
        if (jumpRule.operators === 'equals' && String(e) === jumpRule.answer) {
          console.log(jumpRule.nextQuestionId)
          // setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);
          setCurrent(parseInt(jumpRule.nextQuestionId)-1);

        }
      }
    }
  };

  return(
  <Form.Item name={['question', "'"+question.id+"'"]}   required={question.required===1} >
    <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
  <ProFormDigit
    width="md"
    placeholder="请输入..."
    name={['question', question.id]}
    style={{Width: 60}}
    onChange={onChange}
    rules={[{required: question.required === 1, message: '必填项'}]}
  />

    <QJumpRules
      surveyId={surveyId}
      question={question}
      generateRandom={generateRandom}
      addRespondent={addRespondent}
      setCurrentNum={setCurrentNum}
      setCurrent={setCurrent}
      value={value}
    />
  </Form.Item>);
};

export default QNumber;
