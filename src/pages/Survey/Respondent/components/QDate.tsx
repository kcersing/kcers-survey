

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, DatePicker,Form,Checkbox } from 'antd';
import {ProFormDatePicker, ProFormRate, ProFormTextArea} from '@ant-design/pro-components';
import dayjs from 'dayjs';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';


const QDate = (props) => {

  const { surveyId, question, generateRandom, addRespondent, setCurrentNum } = props;
  const [value, setValue] = useState(0);
  if (!question ){return null}
  const onChange = (date: dayjs.Dayjs | null) => {

    const formattedDate = date?.format('YYYY-MM-DD') || '';
    console.log(formattedDate);
    setValue(e.target.value);
    addRespondent({
      surveyId:surveyId,
      questionId:question.id,
      type:question.type,
      value: [formattedDate],
      sn:generateRandom,
    })
    if (question.jumpRules) {
      for (const jumpRule of question.jumpRules) {
        if (jumpRule.operators === 'equals' && String(formattedDate) === jumpRule.answer) {
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
      format={"YYYY-MM-DD"}
      rules={[{required: question.required === 1, message: '必填项'}]}/>

      <QJumpRules
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        value={value}
      />
    </Form.Item>
  );
};

export default QDate;
