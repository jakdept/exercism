#!/usr/bin/env python3.4
""" Provides a Clock object to be used to track the time"""
# clock.py by Jack Hayhurst
#


class Clock:
    """ Clock is used to track the time, and allows minutes to be added or removed"""

    def __init__(self, hour=0, minute=0):
        self._hour = hour
        self._minute = minute
        self.clean_time()

    def clean_time(self):
        """cleanTime brings the hours and minutes on the clock into proper bounds"""
        while self._minute >= 60:
            self._minute -= 60
            self._hour += 1
        while self._minute < 0:
            self._minute += 60
            self._hour -= 1
        while self._hour < 0:
            self._hour += 24
        while self._hour >= 24:
            self._hour -= 24

    def __str__(self):
        return "{:02d}".format(self._hour) + ":" + '{:02d}'.format(self._minute)

    def add(self, minute):
        """add allows you to add or subtract minutes to/from the clock"""
        self._minute += minute
        self.clean_time()
        return str(self)

    def __eq__(self, other):
        if isinstance(other, self.__class__):
            return self.__dict__ == other.__dict__
        return False