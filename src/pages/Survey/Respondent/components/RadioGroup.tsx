// import {ProFormRadio} from "@ant-design/pro-components";
// import React from "react";
// import { Input } from 'antd';
//
//
//
//
//
//
//
// import React, { useState } from 'react';
// import type { RadioChangeEvent } from 'antd';
// import { Input, Radio } from 'antd';
//
// const style: React.CSSProperties = {
//   display: 'flex',
//   flexDirection: 'column',
//   gap: 8,
// };
//
// const RadioGroup: React.FC = (question:API.Questions,findTargetQuestionIndex,setCurrentNum) => {
//   const [value, setValue] = useState(1);
//
//   const onChange = (e: RadioChangeEvent) => {
//     setValue(e.target.value);
//     // addRespondent({
//     //   surveyId:surveyId,
//     //   type:question.type,
//     //   questionId:question.id,
//     //   value:e.target.value,
//     //   sn:generateRandom,
//     // })
//     if(question.jumpRules){
//       for (const jumpRule of question.jumpRules) {
//         if (jumpRule.operators === 'equals' && String(e.target.value) === jumpRule.answer) {
//           // 找到目标问题在根问题中的索引
//           console.log(jumpRule)
//           console.log(e.target.value)
//
//           const targetIndex = findTargetQuestionIndex(
//             jumpRule.nextQuestionId
//           );
//           console.log("targetIndex",targetIndex)
//
//           if (targetIndex !== -1) {
//             setCurrentNum(targetIndex);
//           }
//         }
//
//       }
//     }
//
//
//
//
//   };
//   let options=[];
//   for (const option of question.options) {
//     if (option.inputs !== 2) {
//       options.push({value: option.content, label: option.content})
//     }else{
//       options.push({
//         value: option.content,
//         label: (
//           <>
//             {option.content}...
//               {value === {option.content}&& (
//                 <Input name={['question', question.id]}  variant="filled" placeholder="其他请输入" style={{width: 120, marginInlineStart: 12}}/>
//               )}
//             </>
//         ),
//       })
//     }
//   }
//
//
//
//
//
//   return (
//     <ProFormRadio.Group
//       style={style}
//       onChange={onChange}
//       value={value}
//       options={options}
//       label={question.content}
//       name={['question', "'"+question.id+"'"]}
//       rules={[{ required: question.required===1, message: '必填项' }]}
//     />
//   );
// };
//
// export default RadioGroup;
