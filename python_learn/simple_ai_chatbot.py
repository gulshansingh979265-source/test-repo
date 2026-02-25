"""
Your first REAL AI chatbot using OpenAI!

This version adds a simple web UI so you can chat
in the browser instead of only in the terminal.
"""

import os

from dotenv import load_dotenv
from flask import Flask, render_template_string, request, redirect, url_for
from openai import OpenAI

# Load environment variables (your API key)
load_dotenv()

# Create OpenAI client
client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))

# In-memory chat history for the web UI
chat_history = []


def ask_ai(question: str) -> str:
    """Ask AI a question and get an answer."""
    # Call OpenAI API
    response = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "user", "content": question}
        ]
    )

    # Extract the answer
    answer = response.choices[0].message.content
    return answer


def main_cli():
    """Original CLI loop (still available if you want it)."""
    print("🤖 AI Assistant Ready!")
    print("Ask me anything! (type 'quit' to exit)\n")

    while True:
        question = input("You: ")

        if question.lower() == "quit":
            print("Goodbye! 👋")
            break

        answer = ask_ai(question)
        print(f"AI: {answer}\n")


# --- Simple Flask web UI ---

app = Flask(__name__)

HTML_TEMPLATE = """
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <title>AI Chatbot</title>
    <style>
      body {
        font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
        margin: 0;
        padding: 0;
        background: #f3f4f6;
        display: flex;
        justify-content: center;
        align-items: stretch;
        height: 100vh;
      }
      .chat-container {
        background: #ffffff;
        width: 100%;
        max-width: 800px;
        margin: 0 auto;
        display: flex;
        flex-direction: column;
        box-shadow: 0 10px 30px rgba(15, 23, 42, 0.2);
      }
      .chat-header {
        padding: 1rem 1.5rem;
        background: #111827;
        color: #f9fafb;
        display: flex;
        align-items: center;
        gap: 0.5rem;
      }
      .chat-header span {
        font-size: 1.4rem;
      }
      .chat-header h1 {
        margin: 0;
        font-size: 1.2rem;
      }
      .chat-messages {
        flex: 1;
        padding: 1rem 1.5rem;
        overflow-y: auto;
        background: #e5e7eb;
      }
      .message {
        max-width: 80%;
        margin-bottom: 0.75rem;
        padding: 0.6rem 0.9rem;
        border-radius: 0.75rem;
        line-height: 1.4;
        white-space: pre-wrap;
      }
      .message.user {
        margin-left: auto;
        background: #2563eb;
        color: #ffffff;
        border-bottom-right-radius: 0.15rem;
      }
      .message.ai {
        margin-right: auto;
        background: #ffffff;
        border-bottom-left-radius: 0.15rem;
        border: 1px solid #d1d5db;
      }
      .chat-input {
        padding: 0.75rem 1rem;
        border-top: 1px solid #e5e7eb;
        background: #f9fafb;
        display: flex;
        gap: 0.5rem;
      }
      .chat-input textarea {
        flex: 1;
        resize: none;
        border-radius: 0.5rem;
        border: 1px solid #d1d5db;
        padding: 0.6rem 0.75rem;
        font-size: 0.95rem;
        font-family: inherit;
      }
      .chat-input button {
        padding: 0 1.1rem;
        border-radius: 0.5rem;
        border: none;
        background: #111827;
        color: #f9fafb;
        font-weight: 500;
        cursor: pointer;
      }
      .chat-input button:hover {
        background: #020617;
      }
    </style>
  </head>
  <body>
    <div class="chat-container">
      <div class="chat-header">
        <span>🤖</span>
        <h1>AI Chatbot</h1>
      </div>
      <div class="chat-messages" id="messages">
        {% if chat_history %}
          {% for role, text in chat_history %}
            <div class="message {{ 'user' if role == 'You' else 'ai' }}">
              {{ text }}
            </div>
          {% endfor %}
        {% else %}
          <div class="message ai">
            Ask me anything about AI, code, or anything else!
          </div>
        {% endif %}
      </div>
      <form class="chat-input" method="post" action="{{ url_for('chatbot') }}">
        <textarea
          name="message"
          rows="2"
          placeholder="Type your message and press Send..."
          required
        ></textarea>
        <button type="submit">Send</button>
      </form>
    </div>
    <script>
      const messagesDiv = document.getElementById("messages");
      if (messagesDiv) {
        messagesDiv.scrollTop = messagesDiv.scrollHeight;
      }
    </script>
  </body>
</html>
"""


@app.route("/chatbot", methods=["GET", "POST"])
def chatbot():
    global chat_history
    if request.method == "POST":
        user_message = (request.form.get("message") or "").strip()
        if user_message:
            chat_history.append(("You", user_message))
            ai_reply = ask_ai(user_message)
            chat_history.append(("AI", ai_reply))
        return redirect(url_for("chatbot"))

    return render_template_string(HTML_TEMPLATE, chat_history=chat_history)


@app.route("/", methods=["GET"])
def root_redirect():
    """Redirect bare / to the chatbot UI path."""
    return redirect(url_for("chatbot"))


def main_web():
    print("Starting web UI at http://127.0.0.1:3000/chatbot ...")
    app.run(host="0.0.0.0", port=3000, debug=True)


if __name__ == "__main__":
    # Default to web UI; change to main_cli() if you prefer terminal.
    main_web()