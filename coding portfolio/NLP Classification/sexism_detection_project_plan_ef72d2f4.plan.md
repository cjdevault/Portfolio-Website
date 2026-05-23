---
name: Sexism Detection Project Plan
overview: Complete the binary sexism detection project by implementing data preprocessing, two feature engineering methods, three different models, comprehensive evaluation, and detailed reporting in the Jupyter notebook.
todos:
  - id: data-loading
    content: Load dataset, explore structure, verify train/test split, and display statistics with markdown explanations
    status: completed
  - id: preprocessing
    content: Implement text preprocessing function (remove URLs, emojis, lowercase, clean text) and apply to train/test sets
    status: completed
  - id: feature-tfidf
    content: Implement TF-IDF feature extraction with appropriate hyperparameters, fit on train only, transform train/test
    status: completed
  - id: feature-bert
    content: Implement BERT embeddings feature extraction using transformers library, encode train/test sets
    status: completed
  - id: model-logreg
    content: Train Logistic Regression with both TF-IDF and BERT features, evaluate on test set
    status: completed
  - id: model-rf
    content: Train Random Forest with both TF-IDF and BERT features, evaluate on test set
    status: completed
  - id: model-svm
    content: Train SVM with both TF-IDF and BERT features, evaluate on test set (may need PCA for BERT)
    status: completed
  - id: evaluation-table
    content: Create comprehensive results DataFrame with all metrics (P, R, F1 for each class + weighted averages) for all 6 combinations
    status: completed
  - id: best-performance
    content: Identify and report best performing model combination with detailed metrics and visualization
    status: completed
  - id: summary-sections
    content: Complete all 6 summary sections (Data Preprocessing, Feature Engineering, Model Selection, Training, Evaluation, AI Usage) with detailed explanations
    status: completed
---

# Sexism Detection Final Project Implementation Plan

## Overview

Implement a complete binary sexism detection system following the EDOS competition requirements. The project requires preprocessing, two feature engineering methods, three models, comprehensive evaluation, and detailed documentation.

## Dataset Analysis

- **File**: `edos_labelled_data.csv` (~5,278 instances)
- **Columns**: `rewire_id`, `text`, `label` (sexist/not sexist), `split` (train/test)
- **Test set**: Already defined in dataset (~1,086 instances based on requirements)
- **Challenge**: Social media text with URLs, emojis, user mentions, informal language

## Implementation Structure

### 1. Data Loading and Exploration

**Location**: First code cell in [`final_project.ipynb`](final_project.ipynb)

- Load CSV and examine structure
- Check label distribution (sexist vs. not sexist)
- Verify train/test split integrity (Rule 1: no test set leakage)
- Display sample texts and statistics
- Add markdown explanation of dataset characteristics

### 2. Data Preprocessing

**Location**: Second code cell with markdown explanation

**Preprocessing steps**:

- Remove URLs and user mentions (`[URL]`, `[USER]`)
- Handle emojis (remove or convert to text)
- Lowercase text
- Remove special characters (keep punctuation for context)
- Handle contractions
- Remove extra whitespace
- Optional: spell correction (if time permits)

**Implementation**:

- Create `preprocess_text()` function
- Apply to train and test sets separately
- Document preprocessing decisions with rationale

### 3. Feature Engineering Method 1: TF-IDF

**Location**: Third code cell with markdown explanation

- Use `TfidfVectorizer` from sklearn
- Parameters to tune:
  - `max_features`: 5000-10000
  - `ngram_range`: (1, 2) for unigrams and bigrams
  - `min_df`: 2-5
  - `max_df`: 0.95
- Fit only on training data
- Transform train and test sets
- Document feature dimensions and rationale

### 4. Feature Engineering Method 2: BERT Embeddings

**Location**: Fourth code cell with markdown explanation

- Use pre-trained BERT model (e.g., `distilbert-base-uncased` via transformers library)
- Generate contextual embeddings for each text
- Options:
  - Use `[CLS]` token embedding (768-dim)
  - Or mean pooling of all token embeddings
- Fit/encode only on training data (for consistency, though BERT is pre-trained)
- Transform train and test sets
- Document model choice and embedding strategy

### 5. Model Implementation

#### Model 1: Logistic Regression

**Location**: Fifth code cell

- Use with both TF-IDF and BERT features
- Hyperparameters: `C`, `penalty`, `solver`
- Train on training set, evaluate on test set

#### Model 2: Random Forest

**Location**: Sixth code cell

- Use with both TF-IDF and BERT features
- Hyperparameters: `n_estimators`, `max_depth`, `min_samples_split`
- Train on training set, evaluate on test set

#### Model 3: Support Vector Machine (SVM)

**Location**: Seventh code cell

- Use with both TF-IDF and BERT features
- Hyperparameters: `C`, `kernel` (linear or RBF)
- Train on training set, evaluate on test set

**Note**: For BERT embeddings, may need dimensionality reduction (PCA) before SVM if memory issues occur.

### 6. Evaluation and Results Table

**Location**: Eighth code cell before "Experimental Results" markdown section

**Evaluation metrics** (using `classification_report` from sklearn):

- Precision, Recall, F1-score for each class (Sexist, Non-Sexist)
- Weighted average Precision, Recall, F1-score

**Results DataFrame**:

- Create pandas DataFrame with columns:
  - Feature + Model (e.g., "TF-IDF + Logistic Regression")
  - Sexist (P), Sexist (R), Sexist (F1)
  - Non-Sexist (P), Non-Sexist (R), Non-Sexist (F1)
  - Weighted (P), Weighted (R), Weighted (F1)
- Minimum 6 rows (2 features × 3 models)
- Display formatted table

### 7. Best Performance Report

**Location**: After results table

- Identify best model (highest weighted F1-score)
- Report:
  - Best feature engineering method
  - Best model
  - Best weighted F1-score
  - Detailed metrics for best model
- Add visualization (optional): confusion matrix for best model

### 8. Project Summary Sections

**Location**: Fill in the existing markdown sections at the end

#### 1. Data Preprocessing

- Document all preprocessing steps
- Explain rationale for each step
- Discuss impact on model performance

#### 2. Feature Engineering

- Compare TF-IDF vs. BERT embeddings
- Discuss advantages/disadvantages
- Explain feature dimensions and sparsity

#### 3. Model Selection and Architecture

- Describe each model (Logistic Regression, Random Forest, SVM)
- Explain why these models were chosen
- Discuss hyperparameters and tuning approach

#### 4. Training and Validation

- Document train/test split (using existing split from dataset)
- Explain evaluation methodology
- Discuss any cross-validation or hyperparameter tuning

#### 5. Evaluation and Results

- Summarize results table
- Analyze which feature+model combinations performed best
- Discuss class-wise performance (sexist vs. non-sexist)
- Interpret weighted F1-scores

#### 6. Use of Generative AI

- Document any AI assistance used
- Specify what was AI-assisted vs. manually implemented
- Be transparent about tool usage

## Technical Requirements

### Libraries Needed

```python
import pandas as pd
import numpy4 as np
from sklearn.model_selection import train_test_split
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.linear_model import LogisticRegression
from sklearn.ensemble import RandomForestClassifier
from sklearn.svm import SVC
from sklearn.metrics import classification_report, confusion_matrix
from transformers import AutoTokenizer, AutoModel
import torch
```

### Key Constraints

- **Rule 1**: Never use test set for training or feature fitting
- **Rule 2**: Document all AI usage
- **Minimum**: 2 feature methods × 3 models = 6 combinations
- **Target**: Weighted F1-score ≥ 0.82 for full points

## File Structure

- All code in [`final_project.ipynb`](final_project.ipynb)
- Dataset: `edos_labelled_data.csv` (already present)
- Output: Results table and summaries in notebook

## Implementation Order

1. Data loading and exploration
2. Preprocessing function and application
3. TF-IDF feature extraction
4. BERT embedding extraction
5. Model training (3 models × 2 features = 6 combinations)
6. Evaluation and results table creation
7. Best performance identification
8. Summary sections completion

## Success Criteria

- ✅ At least 6 model combinations evaluated
- ✅ Results table with all required metrics
- ✅ Best performance clearly identified
- ✅ All summary sections completed
- ✅ No test set leakage
- ✅ Clear explanations throughout