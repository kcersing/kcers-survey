

import React, { useState } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form,Checkbox } from 'antd';
import { ProFormRate, ProFormTextArea } from '@ant-design/pro-components';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import { Rate } from 'antd';

const QRate = (props) => {

  const { surveyId, question, generateRandom, addRespondent, setCurrentNum ,setCurrent} = props;
  const [value, setValue] = useState(0);
  if (!question ){return null}
  const onChange = (e: RadioChangeEvent) => {


    addRespondent({
      surveyId:surveyId,
      questionId:question.id,
      type:question.type,
      value:[e.toString()],
      sn:generateRandom,
    })
    if (question.jumpRules) {
      for (const jumpRule of question.jumpRules) {
        if (jumpRule.operators === 'equals' && String(e) === jumpRule.answer) {
          // setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);
          setCurrent(parseInt(jumpRule.nextQuestionId));
        }
      }
    }

  };


  const character1 = ({ index = 0 }) => {
  if (question.options[0].serial == 0 ){
      return ( index )
    }else {
      return ( index + question.options[0].serial)
    }
  };

  return (<>
      <h3>{question.serial?question.serial+"-":""}{question.content}</h3>
    <Form.Item name={['question', "'"+question.id+"'"]}   >
      <ProFormRate
        style={{color: "rgba(150, 205 ,2050,06)"}}
        fieldProps={{character:character1,allowHalf: false ,count: question.options[0].inputs,style:{fontSize: '30px'} }}

        name={['question', question.id]}
        onChange={onChange}

        rules={[{required: question.required === 1, message: '必填项'}]}
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

export default QRate;
