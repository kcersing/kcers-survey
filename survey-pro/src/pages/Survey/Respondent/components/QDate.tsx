import React, { useEffect, useState } from 'react';
import { DatePicker, Form } from 'antd';
import dayjs from 'dayjs';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QDate = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const defaultDate = dayjs('1955-01-01');
  const [value, setValue] = useState('1955-01-01');
  const form = Form.useFormInstance();
  if (!question) return null;

  const fieldName = ['question', "'" + question.id + "'"];

  useEffect(() => {
    form.setFieldValue(fieldName, defaultDate);
    addRespondent({
      surveyId,
      questionId: question.id,
      type: question.type,
      value: ['1955-01-01'],
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
        disabledDate={(d) => d && d.isAfter('1968-01-01', 'day')}
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
