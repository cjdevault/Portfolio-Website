# Large-Scale NLP Classification & Feature Engineering Pipeline

Machine learning pipeline for large-scale NLP classification using TF-IDF vectorization, BERT embeddings, dimensionality reduction, and supervised learning models to analyze social media text data.

---

## Overview

This project explores natural language processing and machine learning techniques for binary text classification on a large labeled social media dataset. The pipeline includes preprocessing, feature engineering, dimensionality reduction, model training, and evaluation using multiple supervised learning approaches.

The project focuses on balancing predictive performance, interpretability, and computational efficiency while working with high-dimensional text embeddings.

---

## Features

- Text preprocessing and cleaning
- TF-IDF vectorization
- Transformer-based BERT embeddings
- PCA dimensionality reduction
- Logistic Regression and SVM classification
- Precision, Recall, and F1-score evaluation
- Confusion matrix analysis
- Comparative model performance analysis

---

## Technologies Used

- Python
- pandas
- NumPy
- scikit-learn
- Matplotlib
- Jupyter Notebook
- BERT / Transformers

---

## Dataset

- SemEval-2023 Task A dataset
- ~20,000 labeled training samples
- ~5,000 held-out testing samples

---

## Results

- Achieved approximately **0.80 weighted F1-score**
- Reduced embedding dimensionality by over **90%** using PCA while maintaining competitive model performance
