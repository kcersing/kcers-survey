

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';

const style: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
};

const MultipleChoice = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum } = props;
  const [value, setValue] = useState(0);
  if (!question ){return null}
  const onChange = (e: RadioChangeEvent) => {

    setValue(e);
    addRespondent({
      surveyId:surveyId,
      type:question.type,
      questionId:question.id,
      value:e,
      sn:generateRandom,
    })
    if(question.jumpRules){
      for (const jumpRule of question.jumpRules) {
        if (jumpRule.operators === 'equals' &&  e.includes(jumpRule.answer)) {
          setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);
        }

      }
    }

  };

  const onChange1 = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    console.log( e.target.value);
    addRespondent({
      surveyId:surveyId,
      type:"input",
      questionId:question.id,
      value:[e.target.value.toString()],
      sn:generateRandom,
    })
  };

  return (
    <Form.Item name={['question', "'"+question.id+"'"]}   required={question.required===1} >
      <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
      <Checkbox.Group
        onChange={onChange}

        style={style}
        options={question.options.map(option => ({
          value:option.content,
          label: option.inputs!==2? option.content:
            <>
              {option.content}...
              {value!==0 && value &&  value.length > 0   && value.includes(option.content )  && (
                <Input
                  onChange={onChange1}
                  variant="filled"
                  placeholder="please input"
                  style={{ width: 120, marginInlineStart: 12 }}
                />
              )}
            </>
          ,
        }))}
      />

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

export default MultipleChoice;
