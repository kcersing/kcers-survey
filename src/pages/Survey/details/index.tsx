import type { ProColumns } from '@ant-design/pro-components';

import { useRequest } from '@umijs/max';
import {Badge, Card, Descriptions, Divider, Menu, Steps,Image} from 'antd';

import { GridContent, PageContainer, RouteContext,ProTable } from '@ant-design/pro-components';
import {FC, useEffect, useState} from 'react';
import React from 'react';

import {getResponse, getResponseAnswers} from "@/services/ant-design-pro/survey";
import useStyles from './style.style';
import {useParams} from "@@/exports";


const Basic: FC = () => {
  const { styles } = useStyles();
  const { id,sn } = useParams();

  const [loading, setLoading] = useState(false);
  const [response, setResponse] = useState<API.Response>({});
  const [responseAnswers, setResponseAnswers] = useState<API.ResponseAnswers>([]);

  const surveyId = id ? parseInt(id) : 0;
  const responseSn = sn ? sn : "";


  useEffect(() => {
    const getResponses = async () => {
      const [responseRes, responseAnswersRes] = await Promise.all([
        getResponse({sn: responseSn}),
        getResponseAnswers({sn: responseSn}),
      ]);
      setResponse(responseRes.data)
      setResponseAnswers(responseAnswersRes.data)
      setLoading(true)
    }
    getResponses();

  }, [surveyId,responseSn]);


  return (
    <PageContainer>
      <Card bordered={false}>
        <Descriptions
          title="信息"
          style={{
            marginBottom: 32,
          }}
        >
          <Descriptions.Item label="编号">{response?response.sn:""}</Descriptions.Item>
          <Descriptions.Item label="受访人">{response?response.respondent:""}</Descriptions.Item>
          <Descriptions.Item label="受访人联系电话">{response?response.respondentPhone:""}</Descriptions.Item>
          <Descriptions.Item label="调研员"> {response?response.researcher:""}</Descriptions.Item>
          <Descriptions.Item label="调研员联系电话">{response?response.researcherPhone:""}</Descriptions.Item>
          <Descriptions.Item label="合照照片">{
            response?response.pic?.map(v=>{
              return  <Image width={20} src={v} preview={{src: v}}/>
            })   :""
          }</Descriptions.Item>

          <Descriptions.Item label="ip">{response?response.ip:""}</Descriptions.Item>
          <Descriptions.Item label="地址">
            {response?response.area:""}
            {response?response.city:""}
            {response?response.district:""}
            {response?response.village:""}
           {response?response.address:""}
          </Descriptions.Item>
          <Descriptions.Item label="地图地址">{response?response.latitude:""}-{response?response.longitude:""}</Descriptions.Item>
          <Descriptions.Item label="回答数量">{response?response.answerCount:""}</Descriptions.Item>
        </Descriptions>
        <Divider
          style={{
            marginBottom: 32,
          }}
        />
        <div className={styles.title}>详情</div>


          <Card
            style={{
              marginBottom: 24,
            }}
          >
              <Steps
                progressDot
                current={responseAnswers.length+1}
                direction="vertical"
                items={

                  responseAnswers.map(item => ({
                    title: item.content,
                    // 如果有子节点，递归转换
                    description:  <Card >
                        <ul>
                          <li>
                            回答内容:{item.answer.map(v=>{ return (<a>{v}；</a>)} )}
                          </li>
                          <li>
                            其他（补充）:{item.answerText?item.answerText:"无"}
                          </li>
                          <li>
                            回答时间:{item.createdAt}
                          </li>
                        </ul>
                    </Card>,
                  }))

              }
              />

      </Card>

      </Card>
    </PageContainer>
  );
};
export default Basic;
