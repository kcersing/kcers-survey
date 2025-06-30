

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Radio } from 'antd';

const style: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
};

const SingleChoice = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum } = props;
  const [value, setValue] = useState(1);
if (!question ){return null}
  const onChange = (e: RadioChangeEvent) => {

    setValue(e.target.value);
    addRespondent({
      surveyId:surveyId,
      type:question.type,
      questionId:question.id,
      value:e.target.value,
      sn:generateRandom,
    })
    if(question.jumpRules){
      for (const jumpRule of question.jumpRules) {
        if (jumpRule.operators === 'equals' && String(e.target.value) === jumpRule.answer) {
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
      value:e.target.value,
      sn:generateRandom,
    })
  };

  return (
  <Form.Item name={['question', "'"+question.id+"'"]} required >
    <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
    <Radio.Group
      onChange={onChange}

      style={style}
        options={question.options.map(option => ({
        value:option.content,
        label: option.inputs!==2? option.content:
          <>
          {option.content}...
            {value === option.content && (
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
</Form.Item>
  );
};

export default SingleChoice;
