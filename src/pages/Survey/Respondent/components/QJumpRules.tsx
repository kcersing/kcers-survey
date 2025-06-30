

import React from 'react';


import QuestuinSun from '@/pages/survey/respondent/components/QuestuinSun';


const QJumpRules = (props) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum ,value} = props;

let dren=[];
  if(question.jumpRules){
    for (const jumpRule of question.jumpRules) {
      if (jumpRule.operators ==='sub' && String(value) === jumpRule.answer){
        for (const child of    question.children) {
          if ( jumpRule.nextQuestionId === child.id ){
            dren.push(   <QuestuinSun
              surveyId={surveyId}
              question={child}
              generateRandom={generateRandom}
              addRespondent={addRespondent}
              setCurrentNum={setCurrentNum}
            />)
          }
        }
      }}
  }

return dren;
};

export default QJumpRules;
