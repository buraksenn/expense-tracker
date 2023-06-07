# expense-tracker
Open-source telegram bot for tracking your expenses. 

Features:
- Sending receipts store them in drive.
- Store expenses with their categories in Postgres. 
- Write expenses to pre-prepared spreadsheet template that automatically calculates expenses. 
- Get monthly receipts when requested
- Get monthly expenses when requested 

Progress:  
- [X] Telegram bot API
- [X] Required client implementations  
- [ ] Worker manager for expense tracker tasks
- [ ] Expense worker for storing expenses in Postgres.
- [ ] Receipt worker for storing receipts in google cloud drive.
- [ ] Prepare spreadsheet template. 
- [ ] Extend expense worker for writing to spreadsheet.
- [ ] Setup CI/CD and prepare a simple documentation.

Next Steps: 
- Get expenses using OCR from receipt images. 
- TBD
