from flask import Flask, render_template, request
app = Flask(__name__)

@app.route('/headers')
def headers():
    headers = request.headers
    return render_template('headers.html', headers=headers)

if __name__ == "__main__":
    port = 5000
    app.run(host='0.0.0.0', port=port)
