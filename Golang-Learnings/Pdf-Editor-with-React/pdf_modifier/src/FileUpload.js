// src/components/FileUpload.js
import React, { useCallback } from 'react';
import { useDropzone } from 'react-dropzone';

const FileUpload = ({ onFileUpload }) => {

  const onDrop = useCallback((acceptedFiles) => {
    const file = acceptedFiles[0];
    const formData = new FormData();
    formData.append('file', file);

    fetch('http://localhost:8080/upload', {
      method: 'POST',
      body: formData,
    })
      .then(response => response.json())
      .then(data => {
        console.log('File uploaded successfully:', data);
        onFileUpload(file);
      })
      .catch(error => {
        console.error('Error uploading file:', error);
      });
  }, [onFileUpload]);

  // const { getRootProps, getInputProps } = useDropzone({ onDrop, accept: 'application/pdf' });
  const { getRootProps, getInputProps } = useDropzone({ onDrop, accept: 'text/plain' });

  return (
    <div {...getRootProps()} style={{ border: '2px dashed #0087F7', padding: '20px', textAlign: 'center' }}>
      <input {...getInputProps()} />
      <p>Drag & drop a PDF file here, or click to select a file</p>
    </div>
  );
};

export default FileUpload;



