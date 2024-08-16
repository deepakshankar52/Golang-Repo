// src/App.js
import React, { useState } from 'react';
import FileUpload from './FileUpload';
import PdfEditor from './PdfEditor';

const App = () => {
  const [file, setFile] = useState(null);
  const [editedPdf, setEditedPdf] = useState(null);

  const handleFileUpload = (file) => {
    setFile(file);
  };

  const handleSave = (pdfBytes) => {
    const blob = new Blob([pdfBytes], { type: 'application/pdf' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'edited.pdf';
    a.click();
    URL.revokeObjectURL(url);
  };

  return (
    <div>
      {!file ? (
        <FileUpload onFileUpload={handleFileUpload} />
      ) : (
        <PdfEditor file={file} onSave={handleSave} />
      )}
    </div>
  );
};

export default App;
