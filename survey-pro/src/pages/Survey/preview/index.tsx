import React, { useEffect, useState } from 'react';
import { getSurvey, listQuestion } from '@/services/ant-design-pro/survey';
import { useParams } from '@@/exports';
import { ProCard, StepsForm, ProFormText } from '@ant-design/pro-components';
import { message } from 'antd';
import QuestuinSun from '@/pages/survey/respondent/components/QuestuinSun';
import './st.css';

const noop = async () => true;

const Preview = () => {
  const [survey, setSurvey] = useState<API.Survey>({});
  const [questions, setQuestions] = useState<API.Questions[]>([]);
  const generateRandom = 'preview';

  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  useEffect(() => {
    loadSurveyAndQuestions();
  }, []);

  const loadSurveyAndQuestions = async () => {
    try {
      if (!surveyId) return;
      const [surveyData, questionsData] = await Promise.all([
        getSurvey({ id: surveyId }),
        listQuestion({ surveyId }),
      ]);
      setSurvey(surveyData.data || {});
      setQuestions(questionsData.data || []);
    } catch (error: any) {
      message.error(error.message || '加载问卷数据失败');
    }
  };

  return (
    <div className="respondent-container">
      <ProCard boxShadow layout="center">
        <h3>{survey.title || null}</h3>
      </ProCard>

      <ProCard style={{ marginBlockStart: 16 }} boxShadow>
        <StepsForm
          stepsProps={{ direction: 'vertical', size: 'small', current: 1 }}
          onFinish={() => Promise.resolve(true)}
          stepsRender={() => null}
          submitter={{
            render: () => null,
          }}
        >
          {/* 受访人信息 */}
          <StepsForm.StepForm key={`key_${survey.id}`}>
            <ProFormText width="md" label="访谈人姓名" name="respondent" />
            <ProFormText width="md" label="联系电话" name="respondent_phone" />
            <ProFormText width="md" label="调研员姓名" name="researcher" />
            <ProFormText width="md" label="联系电话" name="researcher_phone" />

            {/* 所有题目 — 复用 QuestuinSun 组件 */}
            {questions.map((question) => (
              <React.Fragment key={question.id}>
                {question.children.map((child) => {
                  if (child.type === 'h2') {
                    return (
                      <ProCard key={child.id} style={{ minHeight: 80, marginTop: 8 }}>
                        <h3>{child.content}</h3>
                      </ProCard>
                    );
                  }
                  if (child.type === 'h3') {
                    return (
                      <ProCard key={child.id} style={{ minHeight: 60 }}>
                        <h4>{child.content}</h4>
                      </ProCard>
                    );
                  }
                  return (
                    <div key={child.id} style={{ marginBottom: 12 }}>
                      <QuestuinSun
                        surveyId={surveyId}
                        question={child}
                        generateRandom={generateRandom}
                        addRespondent={noop}
                        setCurrentNum={() => {}}
                        setCurrent={() => {}}
                      />
                    </div>
                  );
                })}
              </React.Fragment>
            ))}
          </StepsForm.StepForm>
        </StepsForm>
      </ProCard>
    </div>
  );
};

export default Preview;
