import board
import busio
import adafruit_ads1x15.ads1115 as ADS
from adafruit_ads1x15.analog_in import AnalogIn
from flask import Flask, jsonify

# Create the I2C bus
i2c = busio.I2C(board.SCL, board.SDA)

# Create the ADC object using the I2C bus
ads = ADS.ADS1115(i2c, address=0x48)

# Create single-ended input on channel 0
chan = AnalogIn(ads, ADS.P0)

# Gain and data rate
ads.gain = 2/3
ads.data_rate = 860

# Flask app for REST API
app = Flask(__name__)

@app.route('/sensor', methods=['GET'])
def get_sensor_data():
    try:
        # Read sensor
        voltage = chan.voltage
        displacement = voltage * 1250 / 5  # Assuming linear relationship between voltage and displacement
        return jsonify({'displacement': displacement, 'voltage': voltage}), 200
    except Exception as e:
        return jsonify({'error': str(e)}), 500


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

