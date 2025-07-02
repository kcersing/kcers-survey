

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox ,message} from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import { ProFormCheckbox }from "@ant-design/pro-components";
const style: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
};

const MultipleChoice = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum,setCurrent } = props;
  const [value, setValue] = useState(0);
  const [disabled, setDisabled] = useState(false);
  if (!question ){return null}
  const onChange = (e: RadioChangeEvent) => {


if (question.id===9 ||question.id===114) {
    if(e.length > 3) {
    message.error("最多选择三个选项");
    // setDisabled(true);
    }
  }
    // else {
    //   setDisabled(false);
    // }
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
          // setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);
          setCurrent(parseInt(jumpRule.nextQuestionId));
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
    <Form.Item name={['question', "'"+question.id+"'"]}  required={question.required===1} >
      <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
      <ProFormCheckbox.Group
        onChange={onChange}
        layout="vertical"
        style={style}
        options={question.options.map(option => ({
          disabled:disabled,
          value:option.content,
          label: option.inputs!==2? option.content:

            <>
              {option.content}...
              {value!==0 && value &&  value.length > 0   && value.includes(option.content )  && (
                <Input
                  onChange={onChange1}
                  variant="filled"
                  placeholder="请输入..."
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
        setCurrent={setCurrent}
        value={value}
      />
    </Form.Item>

  );
};

export default MultipleChoice;
