import React, { useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from '@@/exports';
import { useLocation } from 'react-router-dom';
import { createRespondent } from '@/services/ant-design-pro/survey';
import type { ProFormInstance } from '@ant-design/pro-components';
import { ProCard, StepsForm } from '@ant-design/pro-components';
import { Button, message, Typography } from 'antd';
import { useGeolocation, useSurveyLoader, useRespondentSN } from './hooks';
import QRespondent from '@/pages/survey/respondent/components/QRespondent';
import QuestuinSun from '@/pages/survey/respondent/components/QuestuinSun';
import SFUpload from '@/pages/survey/respondent/components/SFUpload';
import './st.css';

const { Paragraph } = Typography;

const Respondent = () => {
  const formRef = useRef<ProFormInstance>();
  const [currentNum, setCurrentNum] = useState(0);

  const navigate = useNavigate();
  const location = useLocation();
  const urlParams = new URLSearchParams(location.search);
  const sn = urlParams.get('sn');

  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  const { latitude, longitude } = useGeolocation();
  const { survey, questions } = useSurveyLoader(surveyId);
  const { generateRandom } = useRespondentSN(sn);

  // 提交经纬度到后端
  useEffect(() => {
    if (latitude && longitude && generateRandom) {
      addRespondent({
        type: 'location',
        surveyId,
        latitude: latitude.toString(),
        longitude: longitude.toString(),
        sn: generateRandom,
      });
    }
  }, [generateRandom, latitude, longitude]);

  const addRespondent = async (fields: Record<string, any>) => {
    try {
      await createRespondent({ ...fields });
      return true;
    } catch {
      message.error('提交失败');
      return false;
    }
  };

  const renderQuestion = (question: API.Questions, parentname: string): React.ReactNode => {
    if (question.show === 1) return null;

    return (
      <>
        <h4>{parentname}</h4>
        <RenderQuestionControl question={question} />
        {question.children?.length > 0 &&
          question.children.map((child) => (
            <React.Fragment key={child.id}>
              {renderQuestion(child, '')}
            </React.Fragment>
          ))}
      </>
    );
  };

  const RenderQuestionControl = ({ question }: { question: API.Questions }) => {
    if (question.type === 'h2') {
      return <h2 style={{ width: 300 }}><b>{question.content}</b></h2>;
    }
    if (question.type === 'h3') {
      return <h3 style={{ width: 300 }}><b>{question.serial}-{question.content}</b></h3>;
    }
    return (
      <QuestuinSun
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={() => {}}
      />
    );
  };

  const respondentForm = () => (
    <QRespondent
      surveyId={surveyId}
      questions={questions}
      generateRandom={generateRandom}
      addRespondent={addRespondent}
      setCurrentNum={setCurrentNum}
      setCurrent={() => {}}
    />
  );

  const renderThankYou = () => (
    <StepsForm.StepForm name={`key_${questions.length + 2}`} key={`key_${questions.length + 2}`}>
      <ProCard className="thank-you-card" bordered={false}>
        <p className="thank-you-icon" />
        <h2>感谢您参与调查！</h2>
        <p>您的反馈对我们非常重要，感谢您的参与。</p>
        <Button type="primary" onClick={() => navigate(`/survey/${surveyId}/respondent`)} className="finish-button">
          完成
        </Button>
        <Button style={{ left: 10 }} type="primary" onClick={() => navigate(`/survey/${surveyId}/response/${generateRandom}`)} className="finish-button">
          点击查看问卷详情
        </Button>
      </ProCard>
    </StepsForm.StepForm>
  );

  return (
    <div className="respondent-container">
      <ProCard boxShadow layout="center">
        <h3>{survey.title || null}</h3>
      </ProCard>

      <ProCard style={{ width: '100%', marginBlockStart: 16 }} boxShadow>
        <StepsForm
          formRef={formRef}
          stepsProps={{ direction: 'vertical', size: 'small', current: 1 }}
          onFinish={() => Promise.resolve(true)}
          stepsRender={() => null}
          formProps={{ validateMessages: { required: '此项为必填项' } }}
          submitter={{
            render: (props) => {
              if (props.step === 0) {
                return (
                  <ProCard style={{ marginBlockStart: 16 }}>
                    <Button style={{ left: 180 }} type="primary" onClick={() => props.onSubmit?.()}>
                      下一步 {'>'}
                    </Button>
                  </ProCard>
                );
              }
              return (
                <ProCard style={{ marginBlockStart: 16 }}>
                  <Button key="prev" onClick={() => props.onPre?.()}>
                    {'<'} 上一题
                  </Button>
                  <Button style={{ left: 80 }} type="primary" onClick={() => props.onSubmit?.()}>
                    下一步 {'>'}
                  </Button>
                </ProCard>
              );
            },
          }}
        >
          {questions.map((question) => (
            <React.Fragment key={question.id}>
              {question.children.map((que) => (
                <StepsForm.StepForm
                  style={{ width: 360 }}
                  name={que.id}
                  title={que.id}
                  key={que.id}
                >
                  {renderQuestion(que, question.content)}
                </StepsForm.StepForm>
              ))}
            </React.Fragment>
          ))}

          {respondentForm()}

          <SFUpload
            surveyId={surveyId}
            questions={questions}
            generateRandom={generateRandom}
            addRespondent={addRespondent}
            setCurrentNum={setCurrentNum}
            setCurrent={() => {}}
          />

          {renderThankYou()}
        </StepsForm>
      </ProCard>

      <ProCard layout="center" style={{ width: '100%', marginBlockStart: 16 }} boxShadow>
        <Paragraph copyable={{ text: `https://survey.367281.com/survey/${surveyId}/respondent?sn=${generateRandom}` }}>
          当前问卷编号：{generateRandom}
        </Paragraph>
      </ProCard>
    </div>
  );
};

export default Respondent;
