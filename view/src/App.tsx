import React, {FC} from 'react';
import './App.less';
import {Layout, message, Upload} from 'antd';
import {InboxOutlined} from '@ant-design/icons';
import {UploadChangeParam} from "antd/es/upload/interface";
import {Content, Header} from "antd/es/layout/layout";

const {Dragger} = Upload;

const props = {
    name: 'file',
    multiple: true,
    action: '/upload',
    onChange(info: UploadChangeParam) {
        const {status} = info.file;
        if (status !== 'uploading') {
            console.log(info.file, info.fileList);
        }
        if (status === 'done') {
            message.success(`${info.file.name} file uploaded successfully.`);
        } else if (status === 'error') {
            message.error(`${info.file.name} file upload failed.`);
        }
    },
};
const App: FC = () => (
    <Layout>
        <Header>
            <div className="logo"/>
            <h1 className="page-title">在线图像文件转换器</h1>
        </Header>
        <Content style={{padding: '0 50px'}}>
            <Dragger {...props}>
                <p className="ant-upload-drag-icon">
                    <InboxOutlined/>
                </p>
                <p className="ant-upload-text">点击或拖拽文件到该区域上传</p>
                <p className="ant-upload-hint">
                    Support for a single or bulk upload. Strictly prohibit from uploading company data or other
                    band files
                </p>
            </Dragger>
        </Content>
    </Layout>
);

export default App;
