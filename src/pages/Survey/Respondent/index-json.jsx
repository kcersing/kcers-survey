import React, { useEffect, useRef, useState } from "react";
import { getSurvey, listQuestion } from "@/services/ant-design-pro/survey";
import { useParams } from "@@/exports";
import { BetaSchemaForm } from '@ant-design/pro-components';
import { message } from 'antd';
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
function getValues() {
    const str = 'example.com';
    const num = 100;
    return [str, num];
}
function toColumns(questions) {
    let questionColumns = [];
    let stepsArr = [];
    for (const question of questions) {
        console.log(question);
        stepsArr.push({
            title: question.content,
        });
        if (question.type === 'single_choice') {
            let qc = [
                {
                    title: question.content,
                    dataIndex: question.id,
                    valueType: 'radio',
                    formItemProps: {
                        rules: [
                            {
                                required: question.required === 1,
                                message: '此项为必填项',
                            },
                        ],
                    },
                    valueEnum: question.options.map(option => ({
                        text: option.content,
                    })),
                    width: 'm',
                    columns: [
                        {
                            title: '标题',
                            dataIndex: 'titl55555e',
                            formItemProps: {
                                rules: [
                                    {
                                        required: true,
                                        message: '此项为必填项',
                                    },
                                ],
                            },
                        },
                    ],
                },
            ];
            console.log(question.options);
            // for (const opt of question.options) {
            //   if (opt.inputs===2){
            //     console.log(opt);
            //     console.log( qc);
            //     qc[0].columns =  [
            //       {
            //         title:'其他',
            //         dataIndex:"A"+ question.id,
            //         valueType: 'Input',
            //         formItemProps: {
            //           rules: [
            //             {
            //               required: question.required === 1,
            //               message: '此项为必填项',
            //             },
            //           ],
            //         },
            //         width:'m',
            //       }]
            //   }
            // }
            questionColumns.push(qc);
        }
        else if (question.type === 'multiple_choice') {
            questionColumns.push([
                {
                    title: question.content,
                    dataIndex: question.id,
                    valueType: 'checkbox',
                    formItemProps: {
                        rules: [
                            {
                                required: true,
                                message: '此项为必填项',
                            },
                        ],
                    },
                    width: 'm',
                    valueEnum: question.options.map(option => ({
                        text: option.content,
                    })),
                },
            ]);
        }
        else if (question.type === 'date') {
            questionColumns.push([
                {
                    title: question.content,
                    dataIndex: question.id,
                    valueType: 'dateTime',
                    initialValue: new Date(),
                    width: 'md',
                },
            ]);
        }
        else if (question.type === 'rate') {
            questionColumns.push([
                {
                    title: question.content,
                    dataIndex: question.id,
                    valueType: 'rate',
                    width: 'md',
                },
            ]);
        }
        else if (question.type === 'number') {
            questionColumns.push([
                {
                    title: question.content,
                    dataIndex: question.id,
                    valueType: 'inputNumber',
                    width: 'md',
                },
            ]);
        }
        else {
            questionColumns.push([{
                    title: question.content,
                    dataIndex: question.id,
                    valueType: 'group',
                }]);
        }
    }
    // switch (question.type) {
    //   case 'single_choice':
    //
    //     questionColumns.push([
    //       {
    //         title: question.title,
    //         dataIndex: question.id,
    //         valueType: 'ProFormRadio',
    //         options: question.options.map(option => ({
    //           value: option.content,
    //           label: option.content,
    //         })),
    //         width:'m',
    //       },
    //     ])
    //
    //     //
    //
    //
    //
    //     // return (
    //     //   <ProFormRadio.Group
    //     //     name={question.id}
    //     //     options={question.options.map(option => ({
    //     //       value:option.content,
    //     //       label: option.content,
    //     //     }))}
    //     //   />
    //     // );
    //
    //   // case 'multiple_choice':
    //   //   return (
    //   //     <ProFormCheckbox.Group
    //   //       name={question.id}
    //   //       options={question.options.map(option => ({
    //   //         value:option.content,
    //   //         label: option.content,
    //   //       }))}
    //   //     />
    //   //   );
    //
    //   // case 'text':
    //   //   return <ProFormTextArea name={question.id} label="名称"  placeholder="请输入内容..." />;
    //   //
    //   // case 'number':
    //   //   return <ProFormDigit placeholder="请输入数字"
    //   //                        name={question.id}
    //   //                        min={1}
    //   //                        max={10} />;
    //   //
    //   // case 'date':
    //   //   return <ProFormDatePicker  name={question.id} placeholder="请选择日期" />;
    //   //
    //   // case 'rate':
    //   //   return (
    //   //     <ProFormRate name={question.id} label="Rate" />
    //   //   );
    //
    //   // case 'uploadImage':
    //   //   return (
    //   //     <ProFormUploadButton  name={question.id} label="Upload" />
    //   //   );
    //   //
    //   // case 'uploadFile':
    //   //   return (
    //   //     <ProFormUploadButton name={question.id} label="Upload" />
    //   //   );
    //   // case 'h2':
    //   //   return (
    //   //     <h3>{question.content}</h3>
    //   //   );
    //   // case 'h3':
    //   //   return (
    //   //     <h4>{question.content}</h4>
    //   //   );
    //
    //   default:
    //     return null;
    // }
    // questionColumns.push([
    //   {
    //     title: '标题',
    //     dataIndex: 'title',
    //     formItemProps: {
    //       rules: [
    //         {
    //           required: true,
    //           message: '此项为必填项',
    //         },
    //       ],
    //     },
    //     width:'m',
    //   },
    // ])
    // console.log(question)
    return [questionColumns, stepsArr];
}
const Respondent = () => {
    const formRef = useRef();
    const [survey, setSurvey] = useState({});
    const [questions, setQuestions] = useState([]);
    const [steps, setSteps] = useState([]);
    const { id } = useParams();
    const surveyId = id ? parseInt(id) : 1;
    // 加载问卷和问题数据
    useEffect(() => {
        loadSurveyAndQuestions();
    }, []);
    // 加载问卷和问题数据
    const loadSurveyAndQuestions = async () => {
        // try {
        if (!surveyId)
            return;
        const [surveyData, questionsData] = await Promise.all([
            getSurvey({ id: surveyId }),
            listQuestion({ surveyId: surveyId }),
        ]);
        setSurvey(surveyData.data || {});
        const flatArray = treeToArray(questionsData.data);
        const [questionColumns, stepsArr] = toColumns(flatArray);
        setQuestions(questionColumns);
        setSteps(stepsArr);
        // } catch (error: any) {
        //   console.error('加载问卷数据失败', error);
        //   message.error(error.message || '加载问卷数据失败');
        // } finally {
        //
        // }
    };
    return (<BetaSchemaForm layoutType="StepsForm" steps={steps} onCurrentChange={(current) => {
            console.log('current: ', current);
        }} formRef={formRef} onFinish={async (values) => {
            return new Promise((resolve) => {
                console.log(values);
                message.success('提交成功');
                setTimeout(() => {
                    resolve(true);
                    formRef.current?.resetFields();
                }, 2000);
            });
        }} columns={questions}/>);
};
export default Respondent;
//# sourceMappingURL=index-json.jsx.map