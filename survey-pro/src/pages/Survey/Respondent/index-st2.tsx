import React, { useEffect, useState, useRef, useCallback } from 'react';
import { useNavigate, useParams } from '@@/exports';
import { useLocation } from 'react-router-dom';
import { createRespondent } from '@/services/ant-design-pro/survey';
import { ProCard, ProForm } from '@ant-design/pro-components';
import { Button, message, Typography, Form } from 'antd';
import { useGeolocation, useSurveyLoader, useRespondentSN, buildSteps } from './hooks';
import QRespondent from '@/pages/survey/respondent/components/QRespondent';
import QuestuinSun from '@/pages/survey/respondent/components/QuestuinSun';
import SFUpload from '@/pages/survey/respondent/components/SFUpload';
import './st.css';

const { Paragraph } = Typography;

const Respondent = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const urlParams = new URLSearchParams(location.search);
  const sn = urlParams.get('sn');

  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  const { latitude, longitude } = useGeolocation();
  const { survey, questions } = useSurveyLoader(surveyId);
  const { generateRandom } = useRespondentSN(sn);

  const steps = buildSteps(questions);
  const [current, setCurrent] = useState(0);

  const addRespondent = async (fields: Record<string, any>) => {
    try {
      await createRespondent({ ...fields });
      return true;
    } catch {
      message.error('提交失败');
      return false;
    }
  };

  useEffect(() => {
    if (latitude && longitude && generateRandom) {
      addRespondent({
        type: 'location', surveyId,
        latitude: latitude.toString(), longitude: longitude.toString(),
        sn: generateRandom,
      });
    }
  }, [generateRandom, latitude, longitude]);

  // 跳题：根据 sort 找到对应步骤索引
  const pendingJump = useRef<number | null>(null);

  const doGoNext = () => {
    if (pendingJump.current !== null) {
      setCurrent(pendingJump.current);
      pendingJump.current = null;
    } else {
      setCurrent((c) => c + 1);
    }
  };

  // 封装在 ProForm 内部以获取表单实例
  const NextButton = ({ text = "下一步 >", isFirst }: { text?: string; isFirst?: boolean }) => {
    const form = Form.useFormInstance();
    const onClick = async () => {
      try { await form.validateFields(); } catch { message.warning('请填写必填项'); return; }
      doGoNext();
    };
    return isFirst
      ? <div style={{ textAlign: 'center', marginBlockStart: 16 }}><Button type="primary" onClick={onClick}>{text}</Button></div>
      : <div style={{ textAlign: 'center', marginBlockStart: 16 }}>
          <Button onClick={() => { pendingJump.current = null; setCurrent((c) => c - 1); }}>{"<"} 上一题</Button>
          <Button type="primary" style={{ marginLeft: 80 }} onClick={onClick}>{text}</Button>
        </div>;
  };

  // 选中选项时记录跳转目标，等点"下一步"时生效
  const jumpToQuestion = useCallback(
    (n: number) => {
      if (n < 0) { pendingJump.current = null; return; }
      let idx = steps.findIndex((s) => s.questions.some((q) => q.id === n));
      if (idx < 0) idx = steps.findIndex((s) => s.questions.some((q) => q.sort === n));
      if (idx < 0 && n > 0 && n <= steps.length) idx = n - 1;
      if (idx >= 0) pendingJump.current = idx + 1;
    },
    [steps],
  );

  const renderQuestion = (question: API.Questions, pn: string): React.ReactNode => {
    if (question.show === 1) return null;
    const hasSubJump = question.jumpRules?.some((r: any) => r.operators === 'sub');
    return (
      <>
        <h4>{pn}</h4>
        <QuestuinSun
          surveyId={surveyId} question={question}
          generateRandom={generateRandom} addRespondent={addRespondent}
          setCurrentNum={() => {}} setCurrent={jumpToQuestion}
        />
        {!hasSubJump && question.children?.length > 0 &&
          question.children.map((c: any) => (
            <React.Fragment key={c.id}>{renderQuestion(c, '')}</React.Fragment>
          ))}
      </>
    );
  };

  const questionSteps = steps.length;
  const uploadStep = questionSteps + 1;
  const totalSteps = questionSteps + 3;

  return (
    <div className="respondent-container">
      <ProCard boxShadow>
        <h3>{survey?.title || '加载中...'}</h3>
      </ProCard>

      <ProCard style={{ width: '100%', marginBlockStart: 16 }} boxShadow>
        {/* Step 0: 受访人信息 */}
        {current === 0 && (
          <ProForm submitter={false}>
            <QRespondent surveyId={surveyId} questions={questions}
              generateRandom={generateRandom} addRespondent={addRespondent}
              setCurrentNum={() => {}} setCurrent={jumpToQuestion} />
            <NextButton isFirst />
          </ProForm>
        )}

        {/* Steps 1..N: 题目 */}
        {current >= 1 && current <= questionSteps && (() => {
          const step = steps[current - 1];
          return (
            <ProForm submitter={false} key={step.name}>
              <div style={{ width: 360 }}>
                {step.h2heading && <h2><b>{step.h2heading.content}</b></h2>}
                {step.h3heading && <h3><b>{step.h3heading.serial}-{step.h3heading.content}</b></h3>}
                {step.questions.map((q) => (
                  <React.Fragment key={q.id}>{renderQuestion(q, '')}</React.Fragment>
                ))}
              </div>
              <NextButton />
            </ProForm>
          );
        })()}

        {/* Upload step */}
        {current === uploadStep && (
          <ProForm submitter={false}>
            <SFUpload surveyId={surveyId} questions={questions}
              generateRandom={generateRandom} addRespondent={addRespondent}
              setCurrentNum={() => {}} setCurrent={jumpToQuestion} />
            <NextButton text="提交" />
          </ProForm>
        )}

        {/* Thank-you */}
        {current === totalSteps - 1 && (
          <div style={{ textAlign: 'center' }}>
            <h2>感谢您参与调查！</h2>
            <p>您的反馈对我们非常重要，感谢您的参与。</p>
            <Button type="primary"
              onClick={() => navigate(`/survey/${surveyId}/response/${generateRandom}`)}>
              点击查看问卷详情
            </Button>
          </div>
        )}
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
