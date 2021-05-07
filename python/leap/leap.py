#
# leap.py by Jack Hayhurst for exercism.io
# returns whether the given year is a leap year

def is_leap_year(year):
    """ return if the given year is a leap year"""
    if year % 400 is 0:
        return True
    if (year % 4 is 0) and (year % 100 is not 0):
        return True
    return False