// src/components/PdfEditor.js
import React, { useState } from 'react';
import { Document, Page } from 'react-pdf';
import { PDFDocument } from 'pdf-lib';

const PdfEditor = ({ file, onSave }) => {
  const [numPages, setNumPages] = useState(null);

  const onDocumentLoadSuccess = ({ numPages }) => {
    setNumPages(numPages);
  };

  const handleSave = async () => {
    const pdfDoc = await PDFDocument.load(await file.arrayBuffer());
    // Modify the PDF as needed here
    const pdfBytes = await pdfDoc.save();
    onSave(pdfBytes);
  };

  return (
    <div>
      <Document file={file} onLoadSuccess={onDocumentLoadSuccess}>
        {Array.from(new Array(numPages), (el, index) => (
          <Page key={`page_${index + 1}`} pageNumber={index + 1} />
        ))}
      </Document>
      <button onClick={handleSave}>Save Changes</button>
    </div>
  );
};

export default PdfEditor;
