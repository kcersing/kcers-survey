import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox ,message,Alert} from 'antd';
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


    //   if (question.id===467 ||question.id===468) {
    //     if(e.length >= 3) {
    //       setDisabled(true);
    //     }
    //   }


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


    const disableds =(content)=>{
        if (value.length >= 3){
          if (question.id===467 ||question.id===468) {
          console.log(content)
          console.log()
          if(value.includes(content)){
            return false
          }
          return  true
        }
    }
      return false
  }

  return (
<>
    <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
    <Form.Item name={['question', "'"+question.id+"'"]} rules={[{ required: (question.required === 1), message: '这是必填项' }]}>

      <ProFormCheckbox.Group
        onChange={onChange}
        layout="vertical"
        style={style}

        min={2}
        options={question.options.map(option => ({
          disabled:disableds(option.content) ,
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


    </Form.Item>

  <QJumpRules
    surveyId={surveyId}
    question={question}
    generateRandom={generateRandom}
    addRespondent={addRespondent}
    setCurrentNum={setCurrentNum}
    setCurrent={setCurrent}
    value={value}
  />
</>
  );
};

export default MultipleChoice;
