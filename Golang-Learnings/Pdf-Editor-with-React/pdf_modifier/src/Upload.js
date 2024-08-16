import React from "react";
import './App.css';

function Upload() {

  const fileInput = document.getElementById('file-input');

  fileInput.addEventListener('change', (event) => {
    const files = fileInput.files;
    // Process the selected files here
  });

  fileInput.addEventListener('dragover', (event) => {
    event.preventDefault();
  });
  
  fileInput.addEventListener('drop', (event) => {
    event.preventDefault();
    const files = event.dataTransfer.files;
    // Process the dropped files here
  });

  return (
      <div className="Upload">
          <input type="file" id="file-input" multiple></input>
          <label for="file-input">Select a file or drag and drop a file here:</label>
      </div>
  );
}

export default Upload;