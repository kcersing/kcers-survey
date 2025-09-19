

import React from 'react';


import QuestuinSun from '@/pages/survey/respondent/components/QuestuinSun';


const QJumpRules = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum ,setCurrent,value} = props;

let dren=[];




  if(question.jumpRules && question.jumpRules.length>0 && value && value.length>0 ){
    for (const jumpRule of question.jumpRules) {

      if ((jumpRule.operators ==='sub' && String(value) === jumpRule.answer) || (jumpRule.operators ==='sub' &&  value.includes(jumpRule.answer))){
        for (const child of  question.children) {

          if ( jumpRule.nextQuestionId === child.id ){
            dren.push(   <QuestuinSun
              surveyId={surveyId}
              question={child}
              generateRandom={generateRandom}
              addRespondent={addRespondent}
              setCurrentNum={setCurrentNum}
              setCurrent={setCurrent}
            />)
          }
        }
      }}
  }

return dren;
};

export default QJumpRules;
