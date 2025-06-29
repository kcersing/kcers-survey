import React, {useEffect, useRef, useState} from "react";
import {getSurvey, listQuestion, createRespondent} from "@/services/ant-design-pro/survey";

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
import RecursionQuestion from "@/pages/survey/respondent/components/RecursionQuestion";

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
  if(treeNodes){
    treeNodes.forEach(traverse);
  }
  return result;
}


const  Respondent=()=>{

  const formRef = useRef<ProFormInstance>();

  const [survey, setSurvey] = useState<API.Survey>({});
  const [questions, setQuestions] =  useState([]);

  const [current, setCurrent] = useState(0);
  const [currentNum, setCurrentNum] = useState(0);

const [generateRandom, setGenerateRandom] = useState("");



  const formMapRef = useRef<
    React.MutableRefObject<ProFormInstance<any> | undefined>[]
  >([]);


  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  // 加载问卷和问题数据
  useEffect(() => {
    loadSurveyAndQuestions();
  }, [surveyId]);
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
    } catch (error: any) {
      console.error('加载问卷数据失败', error);
      message.error(error.message || '加载问卷数据失败');
    } finally {

    }
  };

  function generateRandomString(length: number): string {
    return Math.random().toString(36).substring(2, 2 + length);
  }

// 查找目标问题的索引
  const findTargetQuestionIndex = (nextQuestionId: string | number) => {
    return questions.findIndex((step) => step.id === nextQuestionId) + 1;
  };



  // useEffect(() => {
  //   console.log("当前问题索引:", currentNum);
  // }, [currentNum]);


  // useEffect(() => {
  //   console.log('formRef:', formRef);
  // }, [formRef]);

  // onChange={(e) => {
  //   console.log(e)
  //   addRespondent({
  //     surveyId: surveyId,
  //     type: "input",
  //     questionId: question.id,
  //     value: e,
  //     sn: generateRandom,
  //   })
  //
  // }}



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
    let opt = [];
    if (question.type === "single_choice") {

      for (const option of question.options) {
        if (option.inputs !== 2) {
          opt.push({value: option.content, label: option.content})
        }else{
          opt.push({
            value: option.content,
            label: (
              <>
                {option.content}...
                <Input name={['question', question.id]} placeholder="其他请输入" style={{width: 120, marginInlineStart: 12}}/>
              </>
            ),
          })
        }
      }
      return (
        <ProFormRadio.Group
          label={question.content}
          options={opt}
          style={style_radio_group}
          name={['question', "'" + question.id + "'"]}
          rules={[{required: question.required === 1, message: '必填项'}]}
          onChange={(e) => {
            addRespondent({
              surveyId:surveyId,
              type:question.type,
              questionId:question.id,
              value:e.target.value,
              sn:generateRandom,
            })

            if (question.jumpRules) {
              for (const jumpRule of question.jumpRules) {
                if (jumpRule.operators === 'equals' && String(e.target.value) === jumpRule.answer) {
                  // 找到目标问题在根问题中的索引

                  const targetIndex = findTargetQuestionIndex(
                    jumpRule.nextQuestionId
                  );

                  if (targetIndex !== -1) {
                    setCurrentNum(targetIndex);
                  }
                }
              }
            }
          }
          }
        />
      );
    }
    if (question.type === 'multiple_choice'){
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
                  name={[question.id, 'input']}
                  placeholder="其他请输入"
                  style={{width: 120, marginInlineStart: 12}}
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
          rules={[{required: question.required === 1, message: '必填项'}]}
          onChange={(e) => {
            addRespondent({
              surveyId:surveyId,
              type:"input",
              questionId:question.id,
              value:e.toString(),
              sn:generateRandom,
            })
            if (question.jumpRules) {
              for (const jumpRule of question.jumpRules) {
                if (
                  jumpRule.operators === "includes" &&
                  e.includes(jumpRule.answer)
                ) {




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
    }
    if (question.type === 'text'){
      return <ProFormTextArea width="md" label={question.content} name={['question', question.id]}
                              placeholder="请输入内容..."
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
                              rules={[{required: question.required === 1, message: '必填项'}]}/>;


    }
    if (question.type === 'number'){
      return <ProFormDigit width="md" label={question.content} placeholder="请输入数字"
                           name={['question', question.id]} style={{maxWidth: 120}}
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
                           rules={[{required: question.required === 1, message: '必填项'}]}/>;

    }
    if (question.type === 'date'){
      return <ProFormDatePicker width="md" label={question.content} name={['question', question.id]}
                                placeholder="请选择日期"

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

                                rules={[{required: question.required === 1, message: '必填项'}]}/>;


    }
    if (question.type === 'rate'){
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

                     rules={[{required: question.required === 1, message: '必填项'}]}/>
      );
    }
    if (question.type === 'uploadImage') {
      return (
        <ProFormUploadButton label={question.content}/>
      )
    }
    if (question.type === 'uploadFile') {
      return (
        <ProFormUploadButton label={question.content}/>
      )
    }
    if (question.type === 'h2') {
      return (<h3> {question.content  } </h3>)
    }
    if (question.type === 'h3') {
      return (<h4> {question.content  } </h4>)
    }


  }




  const style_radio_group: React.CSSProperties = {
    display: 'flex',
    flexDirection: 'column',
    gap: 8,
  };


  const respondent=()=>{
      return (
        <StepsForm.StepForm


          onFinish={(values) => {
            console.log(values)
          }}

          onFieldsChange ={(values) => {
console.log(values)
            //    addRespondent({
            //                             surveyId:surveyId,
            //                             type:"respondent",
            //                             value:e,
            //                             sn:generateRandom,
            //                           })

            // console.log(values)
              }}
          // style={{ minHeight: 300 }}
          name={`key_${questions.length+1}`}
          key={`key_${questions.length+1}`}
          // onBlur={e => {   console.log(e.target.value)}}
        >
          <ProFormText width="md" label="访谈人姓名" rules={[{ required: true, message: '必填项' }]} name={'respondent'} />

          <ProFormText width="md" label="联系电话" rules={[{ required: true, message: '必填项' }]} name={'respondentPhone'} />

          <ProFormText width="md" label="调研员姓名" rules={[{ required: true, message: '必填项' }]} name={'researcher'} />

          <ProFormText width="md" label="联系电话"   rules={[{ required: true, message: '必填项' }]} name={'researcherPhone'} />
          <ProFormText  width="md" label="地图插件预留" name={'ditu'} />
        </StepsForm.StepForm>
      );
  }



  return (
    <div className="respondent-container">
    <ProCard
      boxShadow layout="center"
    >
      <h3>{survey.title ? survey.title  : null}</h3>
      {/*<List items={questions}/>*/}
    </ProCard>
      <ProCard
        style={{ marginBlockStart: 16 }}
        boxShadow
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
      current={current}
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

          return (
            <div>
            <Button key="gotoTwo" onClick={moveToPreviousQuestion}>{'<'} 上一题</Button>

            {currentNum < questions.length  ? (
                <Button type="primary" onClick={moveToNextQuestion}>下一步  {'>'}</Button>
              ) : (
                <Button type="primary" onClick={handleSubmit}>提交 √</Button>
              )}
            </div>);
        },
      }}
    >



          {questions.map((question) => (
            <>
              {question.children.map((que) => (
                <StepsForm.StepForm
                  name={que.content}
                  title={que.content}
                  key={que.content}
                  onFinish={(values) => {
                    console.log(values)
                  }}
                  onValuesChange={(values) => {
                    console.log(values)
                  }}

                  // style={{ minHeight: 300 }}

                >
                  {rq(que,question.content)}

                </StepsForm.StepForm>

              ))}
            </>
          ))}

      {respondent()}

    </StepsForm>

    </ProCard>
    </div>
  )

}
export default Respondent;
