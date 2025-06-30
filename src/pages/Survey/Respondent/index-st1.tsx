import React, {useEffect, useRef, useState} from "react";
import {getSurvey, listQuestion, createRespondent} from "@/services/ant-design-pro/survey";

import {useNavigate, useParams} from "@@/exports";


import {
  ProFormCheckbox,
  ProFormDatePicker,
  ProFormDigit,
  ProFormRadio,
  ProCard,
  ProFormRate,
  ProFormTextArea,
  ProFormUploadButton,
  ProFormText,
  StepsForm,
} from '@ant-design/pro-components';
import { message,Button, Checkbox,Radio, Input} from 'antd';
import type { ProFormInstance } from '@ant-design/pro-components';

import './st.css';
import MultipleChoice from "@/pages/survey/respondent/components/MultipleChoice";
import SingleChoice from '@/pages/survey/respondent/components/SingleChoice';
import QText from '@/pages/survey/respondent/components/QText';
import QNumber from "@/pages/survey/respondent/components/QNumber";
import QRate from '@/pages/survey/respondent/components/QRate';
import QDate from '@/pages/survey/respondent/components/QDate';

import QRespondent from '@/pages/survey/respondent/components/QRespondent';
import QuestuinSun from "@/pages/survey/respondent/components/QuestuinSun";


const  Respondent=()=>{
  const formRef = useRef<ProFormInstance>();
  const [survey, setSurvey] = useState<API.Survey>({});
  const [questions, setQuestions] =  useState([]);
  const [current, setCurrent] = useState(0);
  const [currentNum, setCurrentNum] = useState(0);
  const [generateRandom, setGenerateRandom] = useState("");
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();
  const formMapRef = useRef<
    React.MutableRefObject<ProFormInstance<any> | undefined>[]
  >([]);


  // 新增经纬度状态
  const [latitude, setLatitude] = useState<number | null>(null);
  const [longitude, setLongitude] = useState<number | null>(null);



  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  // 加载问卷和问题数据
  useEffect(() => {
    loadSurveyAndQuestions();

    // 获取经纬度信息
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
          (position) => {
            setLatitude(position.coords.latitude);
            setLongitude(position.coords.longitude);
          },
          (error) => {
            console.error('获取地理位置失败:', error);
            message.error('获取地理位置失败，请检查您的设置');
          }
      );
    } else {
      message.error('您的浏览器不支持地理位置功能');
    }

  }, [surveyId]);


  useEffect(() => {
    if (latitude && longitude) {
      addRespondent({
        type: 'location',
        surveyId:surveyId,
        latitude: latitude.toString(),
        longitude: longitude.toString(),
        sn:generateRandom,
      })
      ;
    }
  }, [generateRandom,latitude,longitude]);



  // 加载问卷和问题数据
  const loadSurveyAndQuestions = async () => {

    try {
      if (!surveyId) return;
      const [surveyData, questionsData] = await Promise.all([
        getSurvey({ id: surveyId }),
        listQuestion({ surveyId: surveyId }),
      ]);
      setSurvey(surveyData.data);
      setQuestions(questionsData.data);
      setGenerateRandom( generateRandomString(18))
      setLoading(true);
    } catch (error: any) {
      console.error('加载问卷数据失败', error);
      message.error(error.message || '加载问卷数据失败');
    } finally {

    }
  };

  function generateRandomString(length: number): string {
    return Math.random().toString(36).substring(2, 2 + length);
  }

  const addRespondent = async (fields) => {

    try {
      await  createRespondent({ ...fields });
      return true;

    } catch (error) {

      message.error('提交失败!');
      return false;
    }
  };

  // 上一步
  const moveToPreviousQuestion = async() => {
    if (current > 0) {
      setCurrent(current - 1);
    }
  };

  // 下一步
  const moveToNextQuestion = async () => {

    // if (formRef.current) {
      try {
        // 验证当前问题
        // formRef.current.validateFields();
        // if (formRef.current) {
        //   const values = await formRef.current.validateFields();
        //   console.log(values)
        //   addRespondent(values)
        // }

// console.log(currentNum)
//         console.log(current)
        console.log(currentNum)
        if (currentNum >0 ) {
          setCurrent(currentNum);
          setCurrentNum(0);
        }else{

            setCurrent(current + 1);

        }

      } catch (error) {
        console.error("当前问题校验失败", error);
        message.error("请填写当前问题的所有必填项");
      }
    // }
  };

  // 提交问卷
  const handleSubmit = async () => {
// console.log(formRef)
    if (formRef.current) {
      try {
        const values = await formRef.current.validateFields();
        console.log('表单提交值:', values);
        formRef.current.submit();
      } catch (error) {
        console.error("表单验证失败", error);
        message.error("请填写所有必填项");
      }
    }
  };

  function List({ items }) {
    return (
      <ul>
        {items.map((item, index) => (
          <li key={index}>
            {item.content} - {item.type}
            {item.children && <List items={item.children} />}
          </li>
        ))}
      </ul>
    );
  }

  const rq =(question,parentname)=> {

    if (question.show === 1) {
      return (<></>);
    }

    return (
      <>
      <h4>{parentname}</h4>

        {RenderQuestionControl(question)}

        {question && question.children && question.children.length > 0 && (
          <>
            {question.children.map((child, index) => (
              <>
                {rq(child,"")}
              </>

            ))}
          </>
        )}


      </>
    )}

  const RenderQuestionControl = (question:API.Questions) => {

    if (question.type === 'h2') {
      return (<h2 style={{width:300}} ><b>{question.content}</b> </h2>)
    }else if (question.type === 'h3') {
      return (<h3 style={{width:300}} ><b>{question.serial}-{question.content }</b> </h3>)
    }else {

      return (
          <QuestuinSun
              surveyId={surveyId}
              question={question}
              generateRandom={generateRandom}
              addRespondent={addRespondent}
              setCurrentNum={setCurrentNum}
          ></QuestuinSun>)
    }

  }

  const respondent=()=>{
      return (
        <QRespondent
          surveyId={surveyId}
          questions={questions}
          generateRandom={generateRandom}
          addRespondent={addRespondent}
          setCurrentNum={setCurrentNum}
        />
      );
  }

  const renderThankYou = () => {
    return (
      <StepsForm.StepForm
        name={`key_${questions.length+2}`}
        key={`key_${questions.length+2}`}
      >
        <ProCard className="thank-you-card" bordered={false}>
          <p className="thank-you-icon" />
          <h2 level={3}>感谢您参与调查！</h2>
          <Text>您的反馈对我们非常重要，我们将根据您的意见改进服务。</Text>
          <Button
            type="primary"
            onClick={() => navigate('/')}
            className="finish-button"
          >
            完成
          </Button>
        </ProCard>
      </StepsForm.StepForm>
    );
  };

  return (
    <div  className="respondent-container">
    <ProCard boxShadow layout="center">
      <h3>{survey.title ? survey.title  : null}</h3>
    </ProCard>
      <ProCard
        style={{ marginBlockStart: 16 }}
        boxShadow
      >
        <StepsForm
            loading={loading}
            formRef={formRef}
            formMapRef={formMapRef}
            stepsProps={{
              direction: 'vertical',
              size:"small",
              current:1,
            }}
            current={current}
            onFinish={(values) => {
              console.log(values);
              return Promise.resolve(true);
            }}
            stepsRender={()=>{
              return (<></>);
            }}

      submitter={{ render: () => {
        return (<>
          <ProCard   style={{ marginBlockStart: 16 }}>
            <Button key="gotoTwo" onClick={moveToPreviousQuestion}>{'<'} 上一题</Button>
            <Button type="primary" onClick={moveToNextQuestion}>下一步  {'>'}</Button>
            {/*{current < questions.length  ? (*/}
            {/*    <Button type="primary" onClick={moveToNextQuestion}>下一步  {'>'}</Button>*/}
            {/*  ) : (*/}
            {/*    <Button type="primary" onClick={handleSubmit}>提交 √</Button>*/}
            {/*  )}*/}
          </ProCard>
        </>);}}}



    >
          {questions.map((question) => (
            <>
              {question.children.map((que) => (
                <StepsForm.StepForm
                  style={{width:360}}
                  name={que.id}
                  title={que.id}
                  key={que.id}
                  onFinish={(values) => {
                    console.log(values)
                  }}
                  onValuesChange={(values) => {
                    console.log(values)
                  }}
                >
                  {rq(que,question.content)}
                </StepsForm.StepForm>
              ))}
            </>
          ))}
      {respondent()}
          {renderThankYou}
    </StepsForm>



    </ProCard>
    </div>
  )

}
export default Respondent;
