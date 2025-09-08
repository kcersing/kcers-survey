import React, {useEffect, useRef, useState} from "react";
import {getSurvey, listQuestion, createRespondent} from "@/services/survey";

import { useParams} from "@@/exports";
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
import { message,Button,  Input} from 'antd';
import type { ProFormInstance } from '@ant-design/pro-components';

import './st.css';
function treeToArray(treeNodes) {
  let result = [];

  //递归函数 traverse，用于处理单个节点
  function traverse(node) {
    const newNode = { ...node };
    delete newNode.children;
    // 将没有子节点的新节点添加到结果数组中
    result.push(newNode);

    // 如果当前节点包含 children 属性（即有子节点）
    if (node.children) {
      node.children.forEach(traverse);
    }
  }

  treeNodes.forEach(traverse);

  return result;
}





const  Preview=()=>{

  const formRef = useRef<ProFormInstance>();

  const [survey, setSurvey] = useState<API.Survey>({});

  const [steps, setSteps] = useState([]);

  const [currentNum, setCurrentNum] = useState(0);

const [generateRandom, setGenerateRandom] = useState("");

  const [questionsh2, setQuestionsh2] = useState("");
  const [questionsh3, setQuestionsh3] = useState("");


  useEffect(() => {
    setGenerateRandom( Math.random().toString(36).substring(2, 2 + length) )
  }, []);

  const formMapRef = useRef<
    React.MutableRefObject<ProFormInstance<any> | undefined>[]
  >([]);


  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  // 加载问卷和问题数据
  useEffect(() => {
    loadSurveyAndQuestions();
  }, []);
  // 加载问卷和问题数据
  const loadSurveyAndQuestions = async () => {

    try {
      if (!surveyId) return;


      const [surveyData, questionsData] = await Promise.all([
        getSurvey({ id: surveyId }),
        listQuestion({ surveyId: surveyId }),
      ]);


      setSurvey(surveyData.data || {});


      const flatArray = treeToArray(questionsData.data);


      setSteps(flatArray);
    } catch (error: any) {
      console.error('加载问卷数据失败', error);
      message.error(error.message || '加载问卷数据失败');
    } finally {

    }

  };


// 查找目标问题的索引
  const findTargetQuestionIndex = (nextQuestionId: string | number) => {
    return steps.findIndex((step) => step.id === nextQuestionId)+1;
  };


  const toPro=(question:API.Questions)=> {
    let opt = [];
    if (!question)return null;



    switch (question.type) {
      case 'single_choice':

        for (const option of question.options) {
          if (option.inputs === 2) {
            opt.push({
              value: option.content,
              label: (
                <>
                  {option.content}...
                  <Input
                    name={['question', question.id]}
                    placeholder="其他请输入"
                    style={{width: 120, marginInlineStart: 12}}


                    onChange={(e) => {
                      console.log(e)

                      addRespondent({
                        surveyId:surveyId,
                        type:"input",
                        questionId:question.id,
                        value:e,
                        sn:generateRandom,
                      })

                    }}
                  />
                </>
              ),
            })
          } else {
            opt.push({value: option.content, label: option.content})
          }
        }
        return (
          <ProFormRadio.Group
            label={question.content}
            options={opt}
            name={['question', "'"+question.id+"'"]}
            rules={[{ required: question.required===1, message: '必填项' }]}
            onChange={(e) => {
console.log(     e,)
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
                    // 找到目标问题在根问题中的索引
                    console.log(jumpRule)
                    console.log(e.target.value)

                    const targetIndex = findTargetQuestionIndex(
                      jumpRule.nextQuestionId
                    );
                    console.log("targetIndex",targetIndex)

                    if (targetIndex !== -1) {
                      setCurrentNum(targetIndex);
                    }
                  }

                }
              }
            }}
          />
        );


      case 'multiple_choice':

        for (const option of question.options) {
          if (option.inputs === 2) {
            opt.push({
              value: option.content,
              label: (
                <>
                  {option.content}...
                  <Input
                    name={['question', question.id]}
                    variant="filled"
                    name={[question.id,'input']}
                    placeholder="其他请输入"
                    style={{ width: 120, marginInlineStart: 12 }}

                    onChange={(e) => {
                      console.log(e)

                      addRespondent({
                        surveyId:surveyId,
                        type:"input",
                        questionId:question.id,
                        value:e,
                        sn:generateRandom,
                      })

                    }}
                  />
                </>
              ),
            });
          } else {
            opt.push({value: option.content, label: option.content})
          }
        }
        return (
          <ProFormCheckbox.Group
            label={question.content}
            options={opt}
            name={['question', question.id]}
            rules={[{ required: question.required===1, message: '必填项' }]}
            onChange={(e) => {

console.log(e)
                addRespondent({
                  surveyId:surveyId,
                  type:question.type,
                  questionId:question.id,
                  value:e,
                  sn:generateRandom,
                })




              if(question.jumpRules){
                for (const jumpRule of question.jumpRules) {
                  if (
                    jumpRule.operators === "includes" &&
                    e.includes(jumpRule.answer)
                  ) {

                    console.log(jumpRule)
                    console.log(e.target.value)

                    const targetIndex = findTargetQuestionIndex(
                      jumpRule.nextQuestionId
                    );
                    if (targetIndex !== -1) {
                      setCurrentNum(targetIndex);
                    }
                  }

                }
              }
            }}
          />
        );
      case 'text':
        return <ProFormTextArea  width="md" label={question.content} name={['question', question.id]} placeholder="请输入内容..."

                                onChange={(e) => {
                                  console.log(e)

                                  addRespondent({
                                    surveyId:surveyId,
                                    type:question.type,
                                    questionId:question.id,
                                    value:e,
                                    sn:generateRandom,
                                  })

                                }}

                                rules={[{ required: question.required===1, message: '必填项' }]}/>;

      case 'number':

        return <ProFormDigit  width="md" label={question.content} placeholder="请输入数字" name={['question', question.id]} style={{ maxWidth: 120 }}

                             onChange={(e) => {
                              console.log(e)

                               addRespondent({
                                 surveyId:surveyId,
                                 questionId:question.id,
                                 type:question.type,
                                 value:e,
                                 sn:generateRandom,
                               })

                             }}

                             rules={[{ required: question.required===1, message: '必填项' }]}/>;
      case 'date':
        return <ProFormDatePicker  width="md" label={question.content}  name={['question', question.id]}
                                  onChange={(e) => {
                                    console.log(e)

                                    addRespondent({
                                      surveyId:surveyId,
                                      questionId:question.id,
                                      type:question.type,
                                      value:e,
                                      sn:generateRandom,
                                    })

                                  }}
                                  placeholder="请选择日期" rules={[{ required: question.required===1, message: '必填项' }]}/>;

      case 'rate':
        return (
          <ProFormRate label={question.content} name={['question', question.id]}
                       onChange={(e) => {
                         console.log(e)

                         addRespondent({
                           surveyId:surveyId,
                           questionId:question.id,
                           type:question.type,
                           value:e,
                           sn:generateRandom,
                         })

                       }}
                       rules={[{ required: question.required===1, message: '必填项' }]} />
        );

      case 'uploadImage':
        return (
          <ProFormUploadButton label={question.content} />
        );

      case 'uploadFile':
        return (
          <ProFormUploadButton label={question.content} />
        );
      case 'h2':
        return (
          <ProCard style={{ minHeight: 240 }}>
            <h3> {question.content}</h3>
          </ProCard>
        );
      case 'h3':
        return (
          <ProCard style={{ minHeight: 240 }}>
           <h4> {question.content}</h4>
          </ProCard>
        );
      default:
        return null;
    }


  }





  useEffect(() => {
    console.log("当前问题索引:", currentNum);
  }, [currentNum]);


  useEffect(() => {
    console.log('formRef:', formRef);
  }, [formRef]);





  const addRespondent = async (fields) => {
      return
    // try {
    //   await  createRespondent({ ...fields });
    //
    //   return true;
    //
    // } catch (error) {
    //
    //   message.error('提交失败!');
    //   return false;
    // }
  };

  // 上一步
  const moveToPreviousQuestion = () => {
    if (currentNum > 0) {
      setCurrentNum(currentNum - 1);
    }
  };


  // 下一步
  const moveToNextQuestion =async () => {
    if (formRef.current) {


      try {
        // 验证当前问题
        formRef.current.validateFields();
        if (currentNum < steps.length ) {
          setCurrentNum(currentNum + 1);
        }
      } catch (error) {
        console.error("当前问题校验失败", error);
        message.error("请填写当前问题的所有必填项");
      }
    }
  };

  // 提交问卷
  const handleSubmit = async () => {

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



  return (
    <div className="respondent-container">
    <ProCard
      boxShadow layout="center"
    >
      <h3>{survey.title ? survey.title  : null}</h3>
    </ProCard>
      <ProCard
        style={{ marginBlockStart: 16 }}
        boxShadow
        title={questionsh2? questionsh2 : null}
      >
    <StepsForm
      formRef={formRef}
      formMapRef={formMapRef}
      stepsProps={{
        direction: 'vertical',
        size:"small",
        current:1,
        // style:{width: 60},
      }}
      current={currentNum}
      onFinish={(values) => {
        // addRespondent(values)
        console.log(values);
        return Promise.resolve(true);
      }}
      stepsRender={()=>{
        return (<></>);
      }}

      submitter={{
        render: (props) => {
          console.log(props)

          return (
            <>
            </>);
        },
      }}
    >

      <StepsForm.StepForm
        key={`key_${survey.id}`}
        // style={{ minHeight: 300 }}
      >

        <ProFormText  width="md" label="访谈人姓名" rules={[{ required: true, message: '必填项' }]} name={'respondent'} />
        <ProFormText  width="md" label="联系电话" rules={[{ required: true, message: '必填项' }]}  name={'respondent_phone'} />
        <ProFormText  width="md" label="调研员姓名" rules={[{ required: true, message: '必填项' }]} name={'researcher'} />
        <ProFormText  width="md" label="联系电话"   rules={[{ required: true, message: '必填项' }]} name={'researcher_phone'} />
        <ProFormText  width="md" label="联系电话" name={'ditu'} />




          {steps.map((question) => (
<div>              {toPro(question)}</div>
            ))}
            </StepsForm.StepForm>
    </StepsForm>

    </ProCard>
    </div>
  );
}

export default Preview;


