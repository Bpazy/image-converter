import React, {FC} from 'react';
import './App.less';
import {message, Upload} from 'antd';
import {InboxOutlined} from '@ant-design/icons';
import {UploadChangeParam} from "antd/lib/upload/interface";

const {Dragger} = Upload;

const props = {
    name: 'file',
    multiple: true,
    action: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
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
);

export default App;
