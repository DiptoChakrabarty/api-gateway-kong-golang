from flask import Flask,jsonify, request

app = Flask(__name__)

@app.route("/api/health")
def health():
    return jsonify(
        {
            "Message": "App is Up and Running python"
        }
    ), 200

@app.route("/api/entry",methods=["POST"])
def entry():
    data = request.get_json()
    name = data["name"]
    email = data["email"]

    return jsonify(
        {
            "Message": f"User {name} with email {email} has been granted access"
        }
    ), 200


if __name__ == "__main__":
    app.run(host="0.0.0.0",debug=True)
