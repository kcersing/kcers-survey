import React, { useState, useEffect } from 'react';
import { DatePicker, Form } from 'antd';
import dayjs from 'dayjs';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const defaultDate = dayjs('1966-08-31');
const defaultDateStr = '1966-08-31';

const QDate = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState(defaultDateStr);
  const form = Form.useFormInstance();
  if (!question) return null;

  const fieldName = ['question', "'" + question.id + "'"];

  useEffect(() => {
    form.setFieldValue(fieldName, defaultDate);
    addRespondent({
      surveyId,
      questionId: question.id,
      type: question.type,
      value: [defaultDateStr],
      sn: generateRandom,
    });
  }, []);

  return (
    <Form.Item
      name={fieldName}
      rules={[{ required: question.required === 1, message: '必填项' }]}
    >
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <DatePicker
        placeholder="请选择日期"
        format="YYYY-MM-DD"
        defaultValue={defaultDate}
        onChange={(date: dayjs.Dayjs | null) => {
          form.setFieldValue(fieldName, date);
          const formattedDate = date?.format('YYYY-MM-DD') || '';
          setValue(formattedDate);
          addRespondent({
            surveyId,
            questionId: question.id,
            type: question.type,
            value: [formattedDate],
            sn: generateRandom,
          });
          handleJump(question, formattedDate, setCurrent);
        }}
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
    </Form.Item>
  );
};

export default QDate;
